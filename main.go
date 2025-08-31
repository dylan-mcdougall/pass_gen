package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	var length int
	var numPasswords int
	var noSymbols bool
	var noNums bool
	var noUpper bool
	var noLower bool
	var includeExtra bool
	var requiredLen int = 0

	flag.IntVar(&length, "l", 16, "Password Length")
	flag.IntVar(&numPasswords, "n", 1, "Number of Passwords")
	flag.BoolVar(&noSymbols, "s", false, "Exclude Symbols")
	flag.BoolVar(&noNums, "d", false, "Exclude Numbers")
	flag.BoolVar(&noUpper, "u", false, "Exclude Uppercase Characters")
	flag.BoolVar(&noLower, "w", false, "Exclude Lowercase Characters")
	flag.BoolVar(&includeExtra, "a", false, "Include Additional Symbols")

	flag.Parse()

	if (noSymbols && noNums && noUpper && noLower && !includeExtra) || length <= 0 {
		fmt.Println("Idk what you expected")
		return
	}

	if !noSymbols {
		requiredLen++
	}
	if !noNums {
		requiredLen++
	}
	if !noUpper {
		requiredLen++
	}
	if !noLower {
		requiredLen++
	}
	if includeExtra {
		requiredLen++
	}

	if length < requiredLen {
		fmt.Printf("Password must be at least %d characters long.", requiredLen)
		return
	}

	allowedChars := charBuilder(noSymbols, noNums, noUpper, noLower, includeExtra)
	charSets := charSetsBuilder(noSymbols, noNums, noUpper, noLower, includeExtra)
	defer cleanMemory(allowedChars)

	for i := 0; i < numPasswords; i++ {
		pass := generatePassword(length, allowedChars, charSets)

		if pass == nil {
			fmt.Println("Error in generation.")
		}

		fmt.Printf("%s\n", string(pass))
		cleanMemory(pass)
	}
}

// Generate a password while ensuring minimum differentiation between char sets
func generatePassword(length int, allowedChars []byte, charSets [][]byte) []byte {
	pass := make([]byte, length)
	allowedCharsLen := int64(len(allowedChars))

	used := make([]bool, length)

	for _, charSet := range charSets {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			cleanMemory(pass)
			return nil
		}

		var position int64
		for {
			pos, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
			if err != nil {
				cleanMemory(pass)
				return nil
			}

			position = pos.Int64()
			if !used[position] {
				break
			}
		}

		pass[position] = charSet[charIndex.Int64()]
		used[position] = true
	}

	for i := 0; i < length; i++ {
		if !used[i] {
			randomIndex, err := rand.Int(rand.Reader, big.NewInt(allowedCharsLen))
			if err != nil {
				cleanMemory(pass)
				return nil
			}
			pass[i] = allowedChars[randomIndex.Int64()]
		}
	}

	return pass
}

// Create sets to guarantee minimum char per available charsets
func charSetsBuilder(noSymbols, noNums, noUpper, noLower, includeExtra bool) [][]byte {
	var sets [][]byte

	if !noLower {
		sets = append(sets, []byte("abcdefghijklmnopqrstuvwxyz"))
	}
	if !noUpper {
		sets = append(sets, []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	}
	if !noNums {
		sets = append(sets, []byte("0123456789"))
	}
	if !noSymbols {
		sets = append(sets, []byte("!@#$%^&*_+-=,.<>?"))
	}
	if includeExtra {
		sets = append(sets, []byte("())[]{}|;:"))
	}

	return sets
}

// Construct available characters
func charBuilder(noSymbols, noNums, noUpper, noLower, includeExtra bool) []byte {
	var chars strings.Builder

	if !noLower {
		chars.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	if !noUpper {
		chars.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if !noNums {
		chars.WriteString("0123456789")
	}
	if !noSymbols {
		chars.WriteString("!@#$%^&*_+-=,.<>?")
	}
	if includeExtra {
		chars.WriteString("())[]{}|;:")
	}

	return []byte(chars.String())
}

func cleanMemory(data []byte) {
	for i := range data {
		data[i] = 0
	}
}
