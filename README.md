# Project Structure Exporter

**Project Structure Exporter** is a Go tool that exports the directory structure of a project into a text file, ignoring directories and files specified in an `export.txt` file (similar to a `.gitignore` file). This is especially useful for documenting or analyzing a project’s structure without including unnecessary or temporary files and directories.

## Features

- Exports the directory structure into a `project_structure.txt` file.
- Ignores directories and files listed in the `export.txt` file.
- Recursively traverses all subdirectories.
- Supports any files and folders excluded in a `.gitignore`-like format.

## Installation

### 1. **Prerequisites**

Make sure Go (at least version 1.16) is installed on your system. If it's not installed, you can download and install it from the [official Go website](https://golang.org/dl/).

### 2. **Clone the project from GitHub**

Clone the repository to your local system:

```bash
git clone https://github.com/JoelKipper/exporter.git
cd exporter
```

### 3. **Initialize the Go module**

If the Go module hasn’t been initialized, run the following command:

```bash
go mod tidy
```

### 4. **Install dependencies**

Ensure all dependencies are installed:

```bash
go get
```

### 5. **Build the program**

Compile the Go program for your operating system:

```bash
go build
```

This will create an executable file that you can use directly.

## Usage

### 1. **Configure the `export.txt` file**

Create a file named export.txt in the same directory as the Go script. In this file, you can specify all directories and files that should be ignored. The format is simple: one line for each file or directory to exclude.

Example `export.txt`:

```txt
.git/
.svn/
node_modules
*.log
*.swp
*.bak
.DS_Store
Thumbs.db
*.exe
```

### 2. **Run the program**
Run the Go program to export the directory structure. The program will recursively scan the current directory and write the structure to a file named `project_structure.txt`.

```bash
./project-structure-exporter
```

### 2. **Result**

After running the program, you will find a `project_structure.txt` file in the same directory. This file contains the project directory structure, with the elements defined in the `export.txt` file excluded.

Example:

```txt
  export.txt
  main.go
  README.md
```

### **Example**

Here is an example of a project structure:

```txt
my_project/
  .git/
  .vscode/
  node_modules/
  src/
    main.go
    utils.go
  README.md
```

With a corresponding `export.txt` file, the exported directory tree would look like this:

```txt
  export.txt
  main.go
  README.md
```

### **Extensions and Customizations**

- Customizing the Ignore List: You can add additional files and directories to the export.txt file to exclude them from the directory tree.
- Support for Additional Formats: Currently, the tool only supports simple directories and file patterns. Future updates may add more filtering options (e.g., regular expressions).

### **License**

This project is licensed under the MIT License – see the LICENSE file for details.

### **Contributing**

Contributions to this project are welcome! If you find bugs or would like to add new features, please open an issue or create a pull request.
