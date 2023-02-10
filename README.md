# File Management (v0.1.1)

[![Go Reference](https://pkg.go.dev/badge/github.com/vieolo/file-management.svg)](https://pkg.go.dev/github.com/vieolo/file-management)

Common functionalities in managing and handling files in GO.

The functions are:

### Check
- `FileExist`: Checks if the file exists or not

### Convert
- `ByteToBase64`: Converts byte array to Base 64

### CSV
- `ReadCsvFile`: Reads a CSV file and converts the contents to an array of arrays of strings

### Download
- `DownloadFileToDisk`: Downloads a file and writes it to the disk
- `DownloadFileToBytes`: Downloads a file and converts it to `[]byte`
- `DownloadFileToBase64`: Downloads a file and converts to base 64

### File System
- `ReadFileToBase64NoError`: Reads a file and convert it to base 64
- `GetDirFilesFlat`: Creates a flat array containing the path to all of the files of a folder and its nested folders.
- `FindFirstFileInFolderContainingText`: Returns the first file in a folder (and its nested folders) that has a text

### Generate
- `CreateDirIfNotExists`: Creates a directory if doesn't exist

### Manipulation
- `ReplaceTextInFile`: Replaces a text in a file
- `InsertTextInFile`: Inserts a text before or after a landmark.
- `AppendToFile`: Adds a text to the bottom of a file

### Movement
- `Copy`: Copies a file
- `CopyDirectory`: Copies a directory and its nested folders
- `CopySymLink`: Copies a sym link