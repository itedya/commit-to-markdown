package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getFileContentsFromCommit(hash string, filePath string) (string, error) {
	cmd := exec.Command("git", "show", hash+":"+filePath)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func getCommitMessage(hash string) (string, error) {
	cmd := exec.Command("git", "show", "-s", "--format=%B", hash)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	for output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}

	return string(output), nil
}

func getChangedFileList(hash string) ([]string, error) {
	cmd := exec.Command("git", "diff-tree", "--no-commit-id", "--name-only", hash, "-r")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	splittedOutput := strings.Split(string(output), "\n")

	for i := 0; i < len(splittedOutput); i++ {
		if splittedOutput[i] == "" {
			splittedOutput = append(splittedOutput[:i], splittedOutput[i+1:]...)
			i--
		}
	}

	return splittedOutput, nil
}

func generateFile(hash string) {
	message, err := getCommitMessage(hash)
	if err != nil {
		fmt.Println("Error getting commit message:", err)
		os.Exit(1)
		return
	}

	files, err := getChangedFileList(hash)
	if err != nil {
		fmt.Println("Error getting changed files:", err)
		os.Exit(1)
		return
	}

	fileContents := make([]string, len(files))

	for i, file := range files {
		content, err := getFileContentsFromCommit(hash, file)
		if err != nil {
			fmt.Println("Error getting file contents:", err)
			os.Exit(1)
			return
		}

		fileContents[i] = content
	}

	fmt.Println("# " + message)
	fmt.Println()
	for i, file := range files {
		fmt.Println("## " + file)
		fmt.Println()
		fmt.Println("```")
		fmt.Println(fileContents[i])
		fmt.Println("```")
		fmt.Println()
	}
}

func printHelp() {
	fmt.Println("Commit to file")
	fmt.Println()
	fmt.Println("     help            - Prints this list")
	fmt.Println("     generate <hash> - Generates markdown file with commit provided by the hash")
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
