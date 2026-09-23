package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main(){

	if len(os.Args) != 2 {
		fmt.Println("Please provide exactly one argument.")
		fmt.Printf("Usage: %s <argument>\n", os.Args[0])
		os.Exit(1)
	}

	// Safely captured 1 arg
	filePath := os.Args[1]
	var file *os.File

	// If the file does not exist create a new one
	_, err := os.Stat(filePath)

	if os.IsNotExist(err){
		fmt.Println("Creating new file")
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	
		_, err = file.WriteString("Hello Distributed Systems World!\n")
		defer file.Close()

		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("File created successfully!")

	} else if err != nil { // Otherwise, if the check did not succeed and the file exists (file stat error)
		fmt.Println("Error determining file status:", err)
		return

	} 

	// Open the file
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	// Safely found file or created file - begin program user input loop
	fmt.Println("<command: read, write, quit> <content to write>")
	var input string
	var command string
	var content string

	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {  

		input = scanner.Text()
		result := strings.SplitN(input, " ", 2)

		command = result[0]
		if len(result) > 1 {
			content = result[1]
		} else {
			content = ""
		}
		
		if command != "quit" && command != "read" && command != "write" {
			fmt.Println("**Invalid command** \n Use one of the following: \n   [read] - display contents of file \n   [write] <arg> - write arg to file \n   [quit] - close program \n**********************")
		}
		
		if command == "quit" {                        // User quits
			fmt.Println("Closing program...")
			return

		} else if command == "read" {                 // User reads
			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
			} else {
				fmt.Println(string(data))
			}

		} else if command == "write" {                // User writes
			_, err = file.WriteString(content)
			if err != nil {
				fmt.Println("Error writing to file:", err)
			}
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Input error:", scanner.Err())
	}
	
}
