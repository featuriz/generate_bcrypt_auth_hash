https://github.com/featuriz/generate_bcrypt_auth_hash

# htpasswd-go 🔐

[![Go Report Card](https://goreportcard.com/badge/github.com/featuriz/generate_bcrypt_auth_hash)](https://goreportcard.com/report/github.com/featuriz/generate_bcrypt_auth_hash)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight, zero-dependency CLI tool written in Go to securely generate and verify `bcrypt` hashes for Basic Authentication (`.htpasswd` files).

Perfect for securing Nginx, Traefik, or Caddy reverse proxies without needing to install the massive `apache2-utils` or `httpd` packages just to get the `htpasswd` binary.

## ✨ Features

- **Zero Dependencies:** Compiles to a single static binary. No Apache, Python, or PHP required.
- **Secure by Default:** Uses modern `bcrypt` hashing, abandoning outdated MD5/SHA1 defaults.
- **Interactive & Scriptable:** Pass arguments directly via CLI for CI/CD automation, or run it without arguments for a secure, hidden terminal prompt that keeps passwords out of your `.bash_history`.
- **Cross-Platform:** Works flawlessly on Linux, macOS, and Windows.

## 🚀 Installation

If you have Go installed, you can easily install the binary globally:

```bash
go install [github.com/featuriz/generate_bcrypt_auth_hash@latest](https://github.com/featuriz/generate_bcrypt_auth_hash@latest)

```

## 📖 Usage

### 1. Generate a Hash

You can pass the username, password, and optional bcrypt cost (4-31, defaults to 12) directly:

```bash
htpasswd-go generate admin mySuperSecretPassword 12
# Output: admin:$2y$12$cl7E...

```

**Interactive Mode:**
Omit the arguments to trigger secure prompts (password input will be hidden):

```bash
$ htpasswd-go generate
Enter username: admin
Enter password:
Enter cost (4-31) [default 12]: 12
admin:$2y$12$cl7E...

```

### 2. Verify a Hash

Check if a password matches an existing bcrypt hash:

```bash
htpasswd-go verify admin mySuperSecretPassword 'admin:$2y$12$cl7E...'
# Output: ✅ Verification SUCCESSFUL for user [admin]. Password matches hash.

```

## 💡 Why was this built?

Managing server infrastructure often requires setting up quick HTTP Basic Auth for staging environments, monitoring dashboards, or custom CMS tools. Installing full web-server utility packages just to generate a single hash is unnecessary bloat.

This tool was originally developed to streamline deployments, keep server footprints small, and secure high-performance web environments at [Featuriz](https://www.google.com/search?q=https://featuriz.in/). We open-sourced it hoping it saves other developers time when configuring their reverse proxies.

## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## 📄 License

[MIT](https://www.google.com/search?q=https://choosealicense.com/licenses/mit/)
