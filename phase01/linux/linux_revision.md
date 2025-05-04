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






