# File Tool
A simple file management tool written in Go.

## Commands
- `create, c`: Create a new file
- `copy, cp`: Copy a file or directory
- `read, r`: Read the contents of a file
- `delete, d`: Delete a file or directory
- `list, ls`: List files and directories

## Usage
- `create, c`: `create <file_name>`
- `copy, cp`: `copy <source_file_or_folder> <destination_file_or_folder>`
- `read, r`: `read <file_name>`
- `delete, d`: `delete <file_or_folder_name>`
- `list, ls`: `list [-r|--recursive] [recursion_limit]`
- `-r, --recursive`: List files and directories recursively
- `recursion_limit`: Limit the recursion depth (default: 2)

## Installation
- Download the latest binary from the [releases page](https://github.com/Branchyz/ft/releases)
- Move the binary to a directory in your PATH
