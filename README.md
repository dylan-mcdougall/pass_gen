# Password Generator

A cryptographically secure password generator written in Go.

## Usage

```bash
go run main.go [options]
```

## Options

| Flag | Description | Default |
|------|-------------|---------|
| `-l` | Password length | 16 |
| `-n` | Number of passwords to generate | 1 |
| `-s` | Exclude symbols (`!@#$%^&*_+-=,.<>?`) | false |
| `-d` | Exclude numbers (`0-9`) | false |
| `-u` | Exclude uppercase letters (`A-Z`) | false |
| `-w` | Exclude lowercase letters (`a-z`) | false |
| `-a` | Include additional symbols (`())[]{}\|;:`) | false |

## Behavior

- Generates cryptographically secure passwords using `crypto/rand`
- Ensures at least one character from each enabled character set
- Securely clears password data from memory after use
- Requires minimum password length based on enabled character sets
- Exits with error if all character sets are disabled

## Examples

```bash
# Generate default 16-character password
go run main.go

# Generate 3 passwords of 20 characters each
go run main.go -l 20 -n 3

# Generate password without symbols
go run main.go -s

# Generate password with only letters and numbers
go run main.go -s -a
```
