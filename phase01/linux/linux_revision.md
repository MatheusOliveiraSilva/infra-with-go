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


