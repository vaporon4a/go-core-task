package main

import (
	"crypto/sha256"
	"fmt"
)

var numDecimal int = 42           // Десятичная система
var numOctal int = 052            // Восьмеричная система
var numHexadecimal int = 0x2A     // Шестнадцатиричная система
var pi float64 = 3.14             // Тип float64
var name string = "Golang"        // Тип string
var isActive bool = true          // Тип bool
var complexNum complex64 = 1 + 2i // Тип complex64

// toString converts any type to string.
// It uses fmt.Sprintf with %v format verb under the hood.
func toString(a any) string {
	return fmt.Sprintf("%v", a)
}

// toRuneSlice converts a string to a slice of runes.
// It is used to satisfy the function hashSHA256 which accepts a slice of runes.
func toRuneSlice(s string) []rune {
	return []rune(s)
}

// makeString takes any number of strings as arguments and returns a concatenated string.
// It is used to combine all the strings that we want to hash together.
func makeString(str ...string) string {
	result := ""
	for _, s := range str {
		result += s
	}
	return result
}

// printType prints the type of the provided argument 'a' to the standard output.
func printType(a any) {
	fmt.Printf("%T\n", a)
}

// hashSHA256 takes a slice of runes, converts it to a byte slice, and returns
// the SHA-256 hash of the byte slice as a hexadecimal string.
func hashSHA256(r []rune) string {
	hasher := sha256.New()
	hasher.Write([]byte(string(r)))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// addSalt adds a salt to a slice of runes to make it more secure.
// It splits the rune slice in half and adds the salt in the middle.
func addSalt(r []rune) []rune {
	salt := []rune("go-2024")
	mid := len(r) / 2
	return append(r[:mid], append(salt, r[mid:]...)...)
}

func main() {
	printType(numDecimal)
	printType(numOctal)
	printType(numHexadecimal)
	printType(pi)
	printType(name)
	printType(isActive)
	printType(complexNum)

	allstring := makeString(toString(numDecimal), toString(numOctal), toString(numHexadecimal), toString(pi), toString(name), toString(isActive), toString(complexNum))

	fmt.Println(hashSHA256(addSalt(toRuneSlice(allstring))))
}
