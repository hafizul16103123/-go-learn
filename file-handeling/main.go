package main

import (
	"io"
	"os"
)
func main()  {



// Create a new file (or truncate if it exists)
os.Create("file.txt") 

// Open an existing file for reading
os.Open("file.txt") 

// Open a file with custom flags (read, write, append, etc.)
os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0644) 

// Write text content to a file
os.WriteFile("file.txt", []byte("Hello, World!"), 0644) 

// Append content to a file
file, _ := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)
file.WriteString("New content") 
file.Close() 

// Read entire file content
data, _ := os.ReadFile("file.txt") 

// Copy file content from source to destination
src, _ := os.Open("source.txt")
dst, _ := os.Create("destination.txt")
io.Copy(dst, src) 
src.Close() 
dst.Close() 

// Move (rename) a file
os.Rename("old.txt", "new.txt") 

// Delete a file
os.Remove("file.txt") 

// Get file size
fileInfo, _ := os.Stat("file.txt")
size := fileInfo.Size() 

// Check if a file exists
_, err := os.Stat("file.txt")
exists := !os.IsNotExist(err) 

// Create a new directory
os.Mkdir("newdir", 0755) 

// Create a directory with parents if they don't exist
os.MkdirAll("parent/child/grandchild", 0755) 

// Delete an empty directory
os.Remove("newdir") 

// Delete a directory and all its contents
os.RemoveAll("parent") 

// List files in a directory
entries, _ := os.ReadDir(".") 
for _, entry := range entries {
    println(entry.Name())
}

// Get current working directory
cwd, _ := os.Getwd() 

// Change the working directory
os.Chdir("newdir") 

}