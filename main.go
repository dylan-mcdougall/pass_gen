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
	var includeUnsafe bool

	flag.IntVar(&length, "length", 16, "Password Length")
	flag.IntVar(&length, "l", 16, "Password Length")
	flag.IntVar(&numPasswords, "number", 1, "Number of Passwords")
	flag.IntVar(&numPasswords, "n", 1, "Number of Passwords")

	flag.BoolVar(&noSymbols, "no-symbols", false, "Exclude Symbols")
	flag.BoolVar(&noNums, "no-num", false, "Exclude Numbers")
	flag.BoolVar(&noUpper, "no-upper", false, "Exclude Uppercase Characters")
	flag.BoolVar(&noLower, "no-lower", false, "Exclude Lowercase Characters")

	flag.BoolVar(&includeUnsafe, "unsafe", false, "Include Problematic Symbols")

	flag.Parse()

	if (noSymbols && noNums && noUpper && noLower && !includeUnsafe) || length <= 0 {
		fmt.Println("Idk what you expected")
		return
	}

	allowedChars := charBuilder(noSymbols, noNums, noUpper, noLower, includeUnsafe)

	allowedCharsLen := int64(len(allowedChars))

	result := make([]string, numPasswords)

	// TODO: Enforce at least one of each available option
	for i := 0; i < numPasswords; i++ {
		res := make([]byte, length)
		for j := 0; j < length; j++ {
			t, err := rand.Int(rand.Reader, big.NewInt(allowedCharsLen))
			if err != nil {
				fmt.Print("Error")
				return
			}

			res[j] = allowedChars[t.Int64()]
		}

		result[i] = string(res)
	}

	for i := range result {
		fmt.Printf("%s\n", result[i])
	}

	// TODO: Move to byte format, compose and clear each byte
	defer func() {
		clear(result)
	}()
}

func charBuilder(noSymbols, noNums, noUpper, noLower, includeUnsafe bool) string {
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
	if includeUnsafe {
		chars.WriteString("())[]{}|;:")
	}

	return chars.String()
}
