# Password Generator

A cryptographically secure password generator written in Go.

## Usage

```bash
go run passgen.go [options]
```

or compile and run
```bash
go build passgen.go
./passgen.go [options]
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
go run passgen.go

# Generate 3 passwords of 20 characters each
go run passgen.go -l 20 -n 3

# Generate password without symbols
go run passgen.go -s

# Generate password with only letters and numbers
go run passgen.go -s -a
```
