package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var ignoredDirs []string

func checkIfIgnored(name string) bool {
	for _, ignored := range ignoredDirs {
		matched, err := path.Match(ignored, name)
		if err == nil && matched {
			return true
		}
		if name == ignored || strings.HasPrefix(name, ignored+"/") {
			return true
		}
	}
	return false
}

func loadIgnoreFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Error when opening the export.txt file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		ignoredDirs = append(ignoredDirs, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error reading the export.txt file: %w", err)
	}

	return nil
}

func main() {
	err := loadIgnoreFile("export.txt")
	if err != nil {
		fmt.Println("Error loading the export.txt file:", err)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error when retrieving the current directory:", err)
		return
	}

	file, err := os.Create(filepath.Join(currentDir, "project_structure.txt"))
	if err != nil {
		fmt.Println("Error when creating the file:", err)
		return
	}
	defer file.Close()

	err = writeDirectoryTree(currentDir, file, 0)
	if err != nil {
		fmt.Println("Error when writing the directory structure:", err)
		return
	}

	fmt.Println("Directory structure was successfully exported to project_structure.txt.")
}

func writeDirectoryTree(dir string, file *os.File, depth int) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("Error when reading the directory %s: %w", dir, err)
	}

	for _, entry := range entries {
		if entry.Name()[0] == '.' || checkIfIgnored(entry.Name()) {
			continue
		}

		for i := 0; i < depth; i++ {
			_, err := file.WriteString("  ")
			if err != nil {
				return fmt.Errorf("Error when writing the indentation: %w", err)
			}
		}

		_, err := file.WriteString(entry.Name() + "\n")
		if err != nil {
			return fmt.Errorf("Error when writing the file name: %w", err)
		}

		if entry.IsDir() {
			err := writeDirectoryTree(filepath.Join(dir, entry.Name()), file, depth+1)
			if err != nil {
				return fmt.Errorf("Error when recursively traversing the directory %s: %w", entry.Name(), err)
			}
		}
	}

	return nil
}
