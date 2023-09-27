package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"
)


// The function `createUser` creates a new user with the specified username, full name, home directory,
// and password, and returns an error if the user already exists or if there are any issues during the
// user creation process.
func createUser(username, fullName, homeDir string, password string) error {
	// Check if the user already exists
	_, err := user.Lookup(username)
	if err == nil {
		return fmt.Errorf("user already exists: %s", username)
	}

	// Create the user
	cmd := exec.Command("sudo", "useradd", "-m", "-d", homeDir, "-c", fullName, username)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Set the user's password
	passwdCmd := exec.Command("sudo", "passwd", username)
	passwdCmd.Stdin = strings.NewReader(password + "\n" + password + "\n")
	passwdCmd.Stdout = os.Stdout
	passwdCmd.Stderr = os.Stderr

	if err := passwdCmd.Run(); err != nil {
		return fmt.Errorf("failed to set password: %v", err)
	}

	fmt.Printf("User %s created successfully.\n", username)
	return nil
}

// The function `promptPassword` prompts the user to enter a password and returns it as a string.
func promptPassword() (string, error) {
	fmt.Print("Enter the user's password: ")
	passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println() // Print a newline after password input

	password := string(passwordBytes)
	return strings.TrimSpace(password), nil
}

func main() {
	username := flag.String("username", "", "Username for the new user")
	fullName := flag.String("full-name", "", "Full name of the new user")
	homeDir := flag.String("home-dir", "/home", "Home directory of the new user")

	flag.Parse()

	if *username == "" || *fullName == "" {
		fmt.Println("Usage: createuser -username <username> -full-name <full name> [-home-dir <home directory>]")
		os.Exit(1)
	}

	password, err := promptPassword()
	if err != nil {
		fmt.Printf("Error reading password: %v\n", err)
		os.Exit(1)
	}

	err = createUser(*username, *fullName, *homeDir, password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
