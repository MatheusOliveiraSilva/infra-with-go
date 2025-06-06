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

### Variables & Globbing

#### Environment Variables
Environment variables are dynamic values that affect the processes or programs on a computer. They can be used to store system-wide values like the location of executables, the default editor, etc.

- **`export MYVAR="value"`**: Sets an environment variable `MYVAR` with the value "value".
- **`echo $MYVAR`**: Retrieves the value of the environment variable `MYVAR`.

#### PATH Manipulation
The `PATH` variable is a critical environment variable that tells the shell where to look for executable files.

- **`export PATH="$HOME/bin:$PATH"`**: Prepends `~/bin` to your `PATH`, allowing executables in `~/bin` to be found before others.

#### Globbing Patterns
Globbing is a way to specify patterns for filenames. It is used in shell commands to match multiple files or directories.

- **`*`**: Matches zero or more characters. Example: `ls *.go` lists all `.go` files.
- **`?`**: Matches exactly one character. Example: `ls file?.txt` matches `file1.txt`, `file2.txt`, etc.
- **`{1..5}`**: Numeric ranges. Example: `echo {1..5}.log` outputs `1.log 2.log ... 5.log`.

These tools are essential for efficient file management and script writing in Unix-like systems.

### Processes

Managing processes is a crucial part of using Unix-like operating systems. Here are some common commands for handling processes:

#### List Running Processes
- **`ps aux`**: Lists all processes by all users.
- **`top`**: Provides an interactive real-time view of running processes.
- **`htop`**: If installed, offers a more user-friendly interface than `top`.

#### Kill a Process by PID
- **`kill 12345`**: Sends a gentle termination signal to the process with PID 12345.
- **`kill -9 12345`**: Forcefully kills the process with PID 12345.

#### Adjust Process Priority
- **`nice -n 10 long_running_command`**: Starts a command with a lower priority.
- **`renice -n -5 -p 12345`**: Raises the priority of the process with PID 12345.

These commands are essential for managing system resources and ensuring efficient operation of processes.

### Permissions

Managing file and directory permissions is crucial for system security and proper access control in Unix-like operating systems.

#### Change File Permissions with `chmod`
The `chmod` command is used to change the permissions of a file or directory. Permissions can be set using either symbolic or numeric modes.

- **Symbolic Mode**: Uses letters to represent permissions.
  - Example: `chmod u+x script.sh` makes the script executable by the owner.
- **Numeric Mode**: Uses octal numbers to set permissions.
  - Example: `chmod 644 file.txt` sets permissions to `rw-r--r--`, allowing the owner to read and write, and others to read only.

#### Understanding Numeric Permissions
File permissions are represented numerically using octal numbers, where each digit represents a different set of permissions.

1. **Permission Types**:
   - **Read (r)**: Allows reading the file or listing the directory contents.
   - **Write (w)**: Allows modifying the file or directory contents.
   - **Execute (x)**: Allows executing the file or accessing the directory.

2. **Permission Groups**:
   - **Owner**: The user who owns the file.
   - **Group**: The group that owns the file.
   - **Others**: All other users.

3. **Numeric Values**:
   - Each permission type is assigned a numeric value:
     - Read (r) = 4
     - Write (w) = 2
     - Execute (x) = 1
   - No permission = 0

4. **Combining Permissions**:
   - Permissions are combined by adding their numeric values:
     - Read + Write = 4 + 2 = 6
     - Read + Execute = 4 + 1 = 5
     - Read + Write + Execute = 4 + 2 + 1 = 7

#### Change Owner or Group with `chown`
The `chown` command is used to change the owner and group of a file or directory.

- **Example**: `sudo chown matheus:staff file.txt` sets the owner to "matheus" and the group to "staff" for `file.txt`.

#### Use `sudo` to Run Commands as Root
The `sudo` command allows a permitted user to run a command as the superuser or another user, as specified by the security policy.

- **Example**: `sudo apt update` runs the `apt update` command as root (example for Debian/Ubuntu).
- **Example**: `sudo -l` lists allowed `sudo` commands for the current user.

These commands are essential for managing access and ensuring that only authorized users can modify or execute files.

### Basic Networking

Understanding basic networking commands is crucial for diagnosing network issues and managing network configurations in Unix-like systems.

#### Check Your IP Addresses and Interfaces
- **`ip a`** (on Linux): Displays all network interfaces and their IP addresses.
- **`ifconfig`** (on macOS or Linux with `net-tools` installed): Shows network interfaces and their configurations, including IP addresses.

#### Test Connectivity
- **`ping google.com`**: Sends ICMP echo requests to test network connectivity to a specified host.
- **`traceroute google.com`**: Displays the route and measures transit delays of packets across an IP network.

#### Inspect Open Ports and Listening Services
- **`ss -tulpn`** (on Linux): Shows open ports and listening services with detailed socket statistics.
- **`lsof -iTCP -sTCP:LISTEN`** (on macOS): Lists TCP ports in the LISTEN state, showing services listening for connections.

#### Fetch a URL and View Headers/Body
- **`curl -I https://api.example.com`**: Fetches HTTP headers of a URL without downloading the body.
- **`curl -v http://localhost:8080`**: Performs a verbose HTTP request, displaying request and response headers and body.

These commands are essential for diagnosing network issues, testing connectivity, and inspecting network configurations and services.