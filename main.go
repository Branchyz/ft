package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "create", "c":
		create()
	case "copy", "cp":
		copy()
	case "read", "r":
		read()
	case "delete", "d":
		delete()
	case "list", "ls":
		list()
	case "help":
		help()
	default:
		fmt.Println("Invalid command, use 'ft help' for usage instructions")
		os.Exit(1)
	}
}

func create() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	fileName := os.Args[2]

	if strings.Contains(fileName, "/") || strings.Contains(fileName, "\\") {
		dir := filepath.Dir(fileName)

		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	_, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func copy() {
	if len(os.Args) < 4 {
		fmt.Println("Please provide a source and destination file name")
		os.Exit(1)
	}

	source := os.Args[2]
	destination := os.Args[3]

	sourceInfo, err := os.Stat(source)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if sourceInfo.IsDir() {
		err = copyDir(source, destination)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		err = copyFile(source, destination)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func copyFile(source, destination string) error {
	data, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	err = os.WriteFile(destination, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func copyDir(source, destination string) error {
	err := os.MkdirAll(destination, 0755)
	if err != nil {
		return err
	}

	files, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, file := range files {
		sourcePath := filepath.Join(source, file.Name())
		destinationPath := filepath.Join(destination, file.Name())

		if file.IsDir() {
			err = copyDir(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(sourcePath, destinationPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func read() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	fileName := os.Args[2]

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		err = readDir(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		err = readFile(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func readFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		return err
	}

	fmt.Printf("Read %d bytes: \n%s\n", count, string(data))

	return nil
}

func readDir(dirName string) error {
	files, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dirName, file.Name())

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			err = readDir(filePath)
			if err != nil {
				return err
			}
		} else {
			err = readFile(filePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func delete() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a file or folder name")
		os.Exit(1)
	}

	fileName := os.Args[2]

	err := os.RemoveAll(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func list() {
	if len(os.Args) > 2 && (os.Args[2] == "-r" || os.Args[2] == "--recursive") {
		if len(os.Args) > 3 {
			recursionLimit, err := strconv.Atoi(os.Args[3])
			if err != nil {
				fmt.Println("Invalid recursion limit")
				os.Exit(1)
			}
			listRecursive(".", 0, recursionLimit)
			return
		} else {
			listRecursive(".", 0, 2)
			return
		}
	}

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("└── DIR  %s\n", file.Name())
		} else {
			fmt.Printf("└── FILE %s\n", file.Name())
		}
	}

}

func listRecursive(dir string, currentLevel int, recursionLimit int) {
	if currentLevel >= recursionLimit {
		return
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s└── DIR  %s\n", strings.Repeat("    ", currentLevel), file.Name())
			listRecursive(dir+"/"+file.Name(), currentLevel+1, recursionLimit)
		} else {
			fmt.Printf("%s└── FILE %s\n", strings.Repeat("    ", currentLevel), file.Name())
		}
	}
}
func help() {
	fmt.Println("Available commands:")
	fmt.Println("create, c - Create a new file")
	fmt.Println("copy, cp - Copy a file or directory")
	fmt.Println("read, r - Read the contents of a file")
	fmt.Println("delete, d - Delete a file or directory")
	fmt.Println("list, ls - List files and directories")

	fmt.Println("\nUsage:")
	fmt.Println("create, c: create <file_name>")
	fmt.Println("copy, cp: copy <source_file_or_folder> <destination_file_or_folder>")
	fmt.Println("read, r: read <file_name>")
	fmt.Println("delete, d: delete <file_or_folder_name>")
	fmt.Println("list, ls: list [-r|--recursive] [recursion_limit]")
	fmt.Println("  -r, --recursive: List files and directories recursively")
	fmt.Println("  recursion_limit: Limit the recursion depth (default: 2)")
}
