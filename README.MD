# devenv

[![Go Build](https://github.com/arashrasoulzadeh/devenv/actions/workflows/go.yml/badge.svg)](https://github.com/arashrasoulzadeh/devenv/actions/workflows/go.yml)

**devenv** is a lightweight CLI tool for managing development environment configuration using simple TOML files.  
It lets you define layered environments (base + overrides) and generate outputs like `.env` files for your projects.

---

## ✨ What It Does

- Reads a TOML configuration file
- Merges **base config + environment overrides**
- Generates structured output files (e.g. `.env`)
- Helps you standardize dev setups across environments (dev, staging, prod, etc.)

---

### 📦 Install

#### macOS (Homebrew)
```sh
brew install arashrasoulzadeh/tap/devenv
```

#### All platforms (manual download)
Head to the [Releases](https://github.com/arashrasoulzadeh/devenv/releases) page and download the appropriate binary for your OS:

- **macOS:** `devenv-macos-intel` or `devenv-macos-arm64`
- **Linux:** `devenv-linux`
- **Windows:** `devenv.exe`

Place the downloaded file in your PATH and make it executable (`chmod +x devenv*` on Linux/macOS).


## 💡 Core Idea

You define:

- A **base configuration**
- One or more **environment overrides** (e.g. `dev`, `prod`)

Then devenv merges them and outputs the final result.

---

## ⚙️ How It Works

### 1. Create a config file (`config.toml`)

```toml
# add your config here

[output]
name = ".env"
type = "dotenv"

[base]
name = "Arash \"The Dev\" Rasoulzadeh"
mysql_host = "127.0.0.1"
mysql_user = "arash"
mysql_pass = "1234"

[dev]
mysql_host = "dev.google.com"
mysql_pass = 5143
```

---

### 2. Run devenv with an environment

```bash
devenv dev
```

---

### 3. Output result

This will generate `.env`:

```env
name=Arash "The Dev" Rasoulzadeh
mysql_host=dev.google.com
mysql_user=arash
mysql_pass=5143
```

---

## 🔁 Merge Rules

1. Start with `[base]`
2. Override with selected environment (e.g. `[dev]`)
3. Flatten into output format

---

## 🚀 Getting Started

### Build from source

```bash
git clone https://github.com/arashrasoulzadeh/devenv.git
cd devenv
go build -o devenv ./...
mv devenv /usr/local/bin/
```

---

### Run

```bash
devenv --env dev
```

---

## 📁 Config Structure

```toml
[output]
name = ".env"
type = "dotenv"

[base]
# default values

[dev]
# overrides
```

---

## ❓ FAQ

**Q: What happens if a key is missing in env?**  
A: It errors! .

**Q: Can I have multiple environments?**  
A: Yes — `[dev]`, `[staging]`, `[prod]`, etc.

---

## 📄 License

MIT
