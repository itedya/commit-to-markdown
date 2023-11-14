package main

import (
	"fmt"
	"os"
	"os/exec"
)

func generateFile(hash string) {
	cmd := exec.Command("git", "show", "-s", "--format=%B", hash)

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}

func printHelp() {
	fmt.Println("Commit to file")	
	fmt.Println()
	fmt.Println("     help     - Prints this list")
	fmt.Println("     generate - Generates markdown file with commit provided by the hash")
	fmt.Println()
}

func main() {
	argc := len(os.Args) - 1
	args := os.Args[1:]

	if argc == 0 {
		fmt.Println("Please provide a command.")
	}

	commandName := args[0]

	if commandName == "help" {
		printHelp()
	} else if commandName == "generate" {
		if argc < 2 {
			fmt.Println("You have to provide a commit hash.")
		}
		
		generateFile(args[1])
	} else {
		fmt.Println("You have to provide a command. See help for list of commands.")
	}
}
