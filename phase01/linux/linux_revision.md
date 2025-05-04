# Linux Revision

### ls

The `ls` command in Unix-like operating systems stands for "list" and is used to list directory contents.

#### `ls -la`
- **`-l` (long format)**: Provides detailed information about each file and directory, including permissions, number of links, owner, group, size, and modification date.
- **`-a` (all files)**: Includes hidden files in the listing.

#### File Permissions
- **Owner**: The user who created the file. Has specific permissions.
- **Group**: A collection of users with specific permissions.
- **Others**: Users not in the owner or group categories.

#### Permission Format
- **`rwxrwxrwx`**: Full access for owner, group, and others.
  - `r`: Read permission
  - `w`: Write permission
  - `x`: Execute permission

This setting allows anyone to read, write, and execute the file, which is generally not recommended for security reasons unless necessary.

#### Example Output
```
matheus@MacBook-Air-de-Matheus infra-with-go % ls -la
total 8
drwxr-xr-x@  5 matheus  staff   160 May  4 15:25 .
drwxr-xr-x   5 matheus  staff   160 May  4 13:55 ..
drwxr-xr-x@ 13 matheus  staff   416 May  4 14:02 .git
-rw-r--r--@  1 matheus  staff  1759 May  4 14:08 README.md
drwxr-xr-x@  3 matheus  staff    96 May  4 15:26 phase01
```

### cd

The `cd` command stands for "change directory" and is used to navigate between directories in a Unix-like operating system.

#### Usage
- **`cd [directory]`**: Changes the current directory to the specified directory.
- **`cd ..`**: Moves up one directory level (to the parent directory).
- **`cd`**: Without any arguments, it returns to the user's home directory.

### pwd

The `pwd` command stands for "print working directory" and is used to display the current directory path in a Unix-like operating system.

#### Usage
- **`pwd`**: Outputs the full path of the current working directory.

### tree

The `tree` command is used to display directories and files in a tree-like format, providing a visual representation of the directory structure.

#### Usage
- **`tree`**: Displays the directory structure of the current directory.

#### Example Output
```
matheus@MacBook-Air-de-Matheus infra-with-go % tree             
.
├── README.md
└── phase01
    └── linux
        └── linux_revision.md

3 directories, 2 files
```

This command is particularly useful for quickly understanding the layout of a project or directory.

### cp

The `cp` command is used to copy files and directories in Unix-like operating systems.

#### Usage
- **`cp [source] [destination]`**: Copies the file or directory from the source path to the destination path.
- **`cp -r [source_directory] [destination_directory]`**: Recursively copies a directory and its contents.

This command is essential for duplicating files and directories, allowing for backup and organization of data.

### mv

The `mv` command is used to move or rename files and directories in Unix-like operating systems.

#### Usage
- **`mv [source] [destination]`**: Moves the file or directory from the source path to the destination path.
- **`mv [old_name] [new_name]`**: Renames a file or directory.

This command is essential for organizing files and directories, allowing for changes in structure and naming.

### rm

The `rm` command is used to remove files and directories in Unix-like operating systems.

#### Usage
- **`rm [file]`**: Removes the specified file.
- **`rm -r [directory]`**: Recursively removes a directory and its contents.
- **`rm -f [file]`**: Forces the removal of a file without prompting for confirmation.

This command is powerful and should be used with caution, as it permanently deletes files and directories.

### mkdir

The `mkdir` command is used to create new directories in Unix-like operating systems.

#### Usage
- **`mkdir [directory_name]`**: Creates a new directory with the specified name.
- **`mkdir -p [path/to/directory]`**: Creates a directory and any necessary parent directories.

This command is useful for organizing files into directories and creating directory structures.

### echo

The `echo` command is used to display a line of text or a variable value, and can also redirect output to files in Unix-like operating systems.

#### Usage
- **`echo [text]`**: Prints the specified text to the terminal.
- **`echo $[variable_name]`**: Displays the value of a variable.
- **`echo [text] > [file]`**: Overwrites the specified file with the text.
- **`echo [text] >> [file]`**: Appends the text to the specified file.

#### Examples
- **`echo "hello" > greeting.txt`**: Overwrites the file `greeting.txt` with the text "hello".
- **`echo "world" >> greeting.txt`**: Appends the text "world" to the file `greeting.txt`.

This command is often used in scripts to output text, variable values, or to manage file contents.

### Pipes & Filters

Pipes (`|`) are used to connect the output of one command directly into the input of another command, allowing for complex command sequences by chaining simple commands together.

#### Common Filters
- **`grep PATTERN`**: Filters lines matching PATTERN.
- **`sort`**: Sorts lines alphabetically or numerically.
- **`uniq`**: Removes duplicate adjacent lines (often used after sort).
- **`wc`**: Counts lines, words, and bytes (use `wc -l` for line count).
- **`head` / `tail`**: Show the first/last N lines.

#### Example Pipeline
```bash
cat access.log | grep 404 | sort | uniq -c | sort -nr | head -n 10
```
- **`cat access.log`**: Outputs the contents of `access.log`.
- **`grep 404`**: Filters lines containing "404".
- **`sort`**: Sorts the filtered lines.
- **`uniq -c`**: Counts and removes duplicate lines, showing the count of each unique line.
- **`sort -nr`**: Sorts the lines numerically in reverse order (most frequent first).
- **`head -n 10`**: Shows the top 10 lines, which are the most frequent "404" entries.

This pipeline is a powerful example of how you can combine simple commands to perform complex data processing tasks.