# Stateless Password Manager in Go
=====================================

## Usage
To use the stateless password manager, follow these steps:

1. **Install Go**: If you haven't already, install Go on your Linux system using your package manager:
	* On Ubuntu/Debian: `sudo apt-get install golang-go`
	* On Fedora/CentOS: `sudo dnf install go`
	* On Arch Linux: `sudo pacman -S go`
2. **Verify Go installation**: Open a terminal and run `go version` to verify that Go is installed correctly.
3. **Run the password manager**: Save the password manager code in a file named `pass.go`, then run it using the command `go run pass.go`.
4. **Enter input**: When prompted, enter the following information:
	* Username
	* Master Password
	* Account Name
5. **Generate password**: The program will generate a password based on the provided input and display it.

## Features
The stateless password manager has the following features:

* Generates passwords based on username, master password, and account name
* Uses SHA-256 hashing for secure password generation
* Applies transformations to ensure passwords meet common requirement rules (e.g., capitalization, special characters)
* Stateless design ensures no passwords are stored

## Why Use a Stateless Password Manager?
----------------------------------------

There are several reasons why you might want to use a stateless password manager:

* **Security**: By not storing any passwords, this manager eliminates the risk of password databases being compromised or leaked.
* **Convenience**: With a stateless password manager, you only need to remember a single master password and a few other details (such as your username and account name) to generate unique, secure passwords for each of your accounts.
* **Flexibility**: This manager can be used on any device, without the need to sync or store any data, making it ideal for use on public computers or other untrusted devices.
* **Compliance**: In some industries or organizations, password managers that store passwords may not be allowed due to security or compliance concerns. A stateless password manager can help meet these requirements.

## Code Explanation
The `generatePassword` function takes three inputs: `username`, `userpassword`, and `account`. It combines these inputs with a fixed entropy string (which can be customized for increased security) and generates a SHA-256 hash. The hashed bytes are then encoded to hex and transformed to create a password that meets common requirement rules.

## Security Notes
* The entropy string can be customized to increase security, but it is essential to back up the custom file and keep it confidential.
* The generated password is not stored anywhere and is only displayed to the user.

## Example Use Case
* Run the program and enter the following input:
	+ Username: `john`
	+ Master Password: `mysecretpassword`
	+ Account Name: `google`
* The program will generate a password based on the provided input and display it.

## Code
The code is written in Go and consists of two main functions: `generatePassword` and `main`. The `generatePassword` function generates the password based on the provided input, while the `main` function handles user input and displays the generated password.

## Customization
To increase security, you can customize the entropy string by replacing the fixed string `"0"` with a random string of your choice. However, make sure to back up the custom file and keep it confidential to avoid losing access to your generated passwords.
