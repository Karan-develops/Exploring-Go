package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 1. Create a new file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 2. Write to the file
	_, err = file.WriteString("Hello, Go Files!\nThis is a test file.")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully.")

	// 3. Open the file for reading
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 4. Read the file using ioutil
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content using ioutil:")
	fmt.Println(string(content))

	// 5. Read the file line by line using bufio
	fmt.Println("File content line by line:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file line by line:", err)
	}

	// 6. Read the file in chunks
	file.Seek(0, io.SeekStart) // Reset file pointer to the beginning
	buffer := make([]byte, 10) // Read in chunks of 10 bytes
	fmt.Println("File content in chunks:")
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file in chunks:", err)
			return
		}
		fmt.Print(string(buffer[:n]))
	}
	fmt.Println()

	// 7. Delete the file
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully.")
}
