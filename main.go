package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func main() {
	// Require at least the command argument
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := strings.ToLower(os.Args[1])
	switch command {
	case "generate", "gen":
		// Pass arguments excluding the binary name and command
		handleGenerate(os.Args[2:])
	case "verify", "check":
		handleVerify(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown command '%s'\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  htpasswd-go generate [username] [password] [cost]")
	fmt.Println("  htpasswd-go verify   [username] [password] [hash]")
	fmt.Println("\nNote: If arguments are omitted, the tool will securely prompt you for them.")
}

// --- GENERATE COMMAND ---
func handleGenerate(args []string) {
	var username, password string
	cost := 12 // Default cost fallback

	// 1. Username
	if len(args) >= 1 {
		username = args[0]
	} else {
		username = promptInput("Enter username: ")
	}

	// 2. Password
	if len(args) >= 2 {
		password = args[1]
	} else {
		password = promptPassword("Enter password: ")
	}

	// 3. Cost (This is the logic that was missing the interactive fallback)
	if len(args) >= 3 {
		// CLI argument provided
		parsedCost, err := strconv.Atoi(args[2])
		if err != nil || parsedCost < bcrypt.MinCost || parsedCost > bcrypt.MaxCost {
			fmt.Fprintf(os.Stderr, "Invalid cost '%s'. Must be an integer between %d and %d.\n", args[2], bcrypt.MinCost, bcrypt.MaxCost)
			os.Exit(1)
		}
		cost = parsedCost
	} else {
		// Interactively prompt for cost
		costInput := promptInputOptional(fmt.Sprintf("Enter cost (%d-%d) [default 12]: ", bcrypt.MinCost, bcrypt.MaxCost))
		if costInput != "" { // If not empty, parse it
			parsedCost, err := strconv.Atoi(costInput)
			if err != nil || parsedCost < bcrypt.MinCost || parsedCost > bcrypt.MaxCost {
				fmt.Fprintf(os.Stderr, "Error: Cost must be an integer between %d and %d.\n", bcrypt.MinCost, bcrypt.MaxCost)
				os.Exit(1)
			}
			cost = parsedCost
		}
	}

	// 4. Generate Hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating hash: %v\n", err)
		os.Exit(1)
	}

	// Output
	fmt.Printf("%s:%s\n", username, hash)
}

// --- VERIFY COMMAND ---
func handleVerify(args []string) {
	var username, password, hash string

	// 1. Username
	if len(args) >= 1 {
		username = args[0]
	} else {
		username = promptInput("Enter username: ")
	}

	// 2. Password
	if len(args) >= 2 {
		password = args[1]
	} else {
		password = promptPassword("Enter password: ")
	}

	// 3. Hash
	if len(args) >= 3 {
		hash = args[2]
	} else {
		hash = promptInput("Enter bcrypt hash (or full htpasswd line): ")
	}

	// Clean hash if it contains the "username:" prefix
	if strings.Contains(hash, ":") {
		parts := strings.SplitN(hash, ":", 2)
		hash = parts[1]
	}

	// 4. Verify
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("❌ Verification FAILED for user [%s]. Invalid password or corrupted hash.\n", username)
		os.Exit(1)
	}

	fmt.Printf("✅ Verification SUCCESSFUL for user [%s]. Password matches hash.\n", username)
}

// --- HELPER CLI PROMPTS ---

// promptInput requires a value to be entered
func promptInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Fprintln(os.Stderr, "Error: Field cannot be empty.")
		os.Exit(1)
	}
	return input
}

// promptInputOptional allows the user to press Enter and return an empty string
func promptInputOptional(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// promptPassword hides terminal echo
func promptPassword(prompt string) string {
	fmt.Print(prompt)
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // Print newline since Enter is hidden
	if err != nil || len(passwordBytes) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Invalid or empty password.")
		os.Exit(1)
	}
	return string(passwordBytes)
}
