package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// generatePassword generates a consistent pseudorandom password based on the provided inputs
func generatePassword(username, userpassword, account string) string {
	entropy := "0" //Adding your own random string here will greatly increase security, but make sure you back up your custom file and don't share it
	data := username + userpassword + account + entropy

	// Create a new SHA-256 hash
	hasher := sha256.New()
	hasher.Write([]byte(data))                                // Write the data to the hasher
	passwordBytes := hasher.Sum(nil)                          // Get the hashed bytes
	encodedPassword := hex.EncodeToString(passwordBytes)[:32] // Encode to hex and take the first 32 characters, you can change this to always have longer or shorter passwords

	if len(encodedPassword) == 0 {
		return "!"
	}

	runes := []rune(encodedPassword)

	// Capitalize the first alphabetic character.
	capitalized := false
	for i, r := range runes {
		if !capitalized && (r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') { // Check if it's an alphabet character
			runes[i] = r - ('a' - 'A') // Convert to uppercase
			capitalized = true         // Mark as capitalized
						  // these are to fit most password requirement rules
		}
	}

	// Replace the last lowercase letter with '!'.
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = '!'
			break
		}
	}

	return string(runes)
}

func main() {
	var username, userpassword, account string
	fmt.Print("Enter Username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter Master Password : ")
	fmt.Scanln(&userpassword)
	fmt.Print("Enter Account Name: ")
	fmt.Scanln(&account)

	password := generatePassword(username, userpassword, account)
	fmt.Printf("Generated Password: %s\n", password)
}


