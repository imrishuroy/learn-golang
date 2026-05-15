package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/*
  FILE I/O IN GO

KEY POINTS:
  - os package: Low-level file operations (Open, Create, Remove)
  - io package: Reader/Writer interfaces
  - bufio package: Buffered I/O for efficiency
  - encoding/json, encoding/csv: Structured data formats

COMMON OPERATIONS:
  os.ReadFile(path)           // Read entire file into []byte
  os.WriteFile(path, data, perm) // Write []byte to file
  os.Open(path)               // Open for reading
  os.Create(path)             // Create/truncate for writing
  os.OpenFile(path, flag, perm) // Open with custom flags

FILE PERMISSIONS (Unix):
  0644 = rw-r--r-- (owner read/write, others read)
  0755 = rwxr-xr-x (executable)
  0600 = rw------- (owner only)

ALWAYS:
  - Check errors after file operations
  - Use defer file.Close() immediately after opening
  - Handle partial reads/writes

*/

// ChargeSession for JSON examples
type ChargeSession struct {
	Id        string `json:"id"`
	Watts     int    `json:"watts"`
	Vin       string `json:"vin"`
	Timestamp string `json:"timestamp"`
}

func main() {
	// 1. WRITING FILES - SIMPLE WAY

	fmt.Println("--- Writing Files (Simple) ---")

	content := []byte("Hello, Go!\nThis is line 2.\nThis is line 3.")
	err := os.WriteFile("example.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Created example.txt")

	// 2. READING FILES - SIMPLE WAY

	fmt.Println("\n--- Reading Files (Simple) ---")

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Contents:")
	fmt.Println(string(data))

	// 3. CHECKING IF FILE EXISTS

	fmt.Println("\n--- File Existence Check ---")

	if _, err := os.Stat("example.txt"); err == nil {
		fmt.Println("example.txt exists")
	} else if os.IsNotExist(err) {
		fmt.Println("example.txt does not exist")
	} else {
		fmt.Println("Error checking file:", err)
	}

	// 4. OPENING FILES WITH DEFER

	fmt.Println("\n--- Opening Files (with defer) ---")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Always close files!

	// Read file info
	info, _ := file.Stat()
	fmt.Printf("File: %s, Size: %d bytes\n", info.Name(), info.Size())

	// 5. BUFFERED READING (LINE BY LINE)

	fmt.Println("\n--- Reading Line by Line ---")

	file2, _ := os.Open("example.txt")
	defer file2.Close()

	scanner := bufio.NewScanner(file2)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning:", err)
	}

	// 6. WRITING WITH BUFIO (EFFICIENT)

	fmt.Println("\n--- Buffered Writing ---")

	file3, err := os.Create("buffered.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file3.Close()

	writer := bufio.NewWriter(file3)
	writer.WriteString("Line 1: Buffered writing is efficient\n")
	writer.WriteString("Line 2: Reduces system calls\n")
	writer.WriteString("Line 3: Good for many small writes\n")
	writer.Flush() // Don't forget to flush!

	fmt.Println("Created buffered.txt")

	// 7. APPENDING TO FILES

	fmt.Println("\n--- Appending to Files ---")

	// O_APPEND: Append mode, O_WRONLY: Write only, O_CREATE: Create if not exists
	file4, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening for append:", err)
		return
	}
	defer file4.Close()

	file4.WriteString("\nAppended line!")
	fmt.Println("Appended to example.txt")

	// 8. COPYING FILES

	fmt.Println("\n--- Copying Files ---")

	err = copyFile("example.txt", "example_copy.txt")
	if err != nil {
		fmt.Println("Error copying:", err)
	} else {
		fmt.Println("Copied to example_copy.txt")
	}

	// 9. WORKING WITH JSON FILES

	fmt.Println("\n--- JSON Files ---")

	// Write JSON
	sessions := []ChargeSession{
		{Id: "CS001", Watts: 420, Vin: "ABC123", Timestamp: "2024-01-01T10:00:00Z"},
		{Id: "CS002", Watts: 350, Vin: "XYZ789", Timestamp: "2024-01-01T11:00:00Z"},
	}

	jsonData, _ := json.MarshalIndent(sessions, "", "  ")
	os.WriteFile("sessions.json", jsonData, 0644)
	fmt.Println("Created sessions.json")

	// Read JSON
	jsonFile, _ := os.ReadFile("sessions.json")
	var loadedSessions []ChargeSession
	json.Unmarshal(jsonFile, &loadedSessions)
	fmt.Println("Loaded sessions:", loadedSessions)

	// 10. WORKING WITH CSV FILES

	fmt.Println("\n--- CSV Files ---")

	// Write CSV
	csvFile, _ := os.Create("data.csv")
	csvWriter := csv.NewWriter(csvFile)

	csvWriter.Write([]string{"ID", "Name", "Score"})
	csvWriter.Write([]string{"1", "Alice", "95"})
	csvWriter.Write([]string{"2", "Bob", "87"})
	csvWriter.Write([]string{"3", "Charlie", "92"})
	csvWriter.Flush()
	csvFile.Close()
	fmt.Println("Created data.csv")

	// Read CSV
	csvFile2, _ := os.Open("data.csv")
	defer csvFile2.Close()

	csvReader := csv.NewReader(csvFile2)
	records, _ := csvReader.ReadAll()

	fmt.Println("CSV contents:")
	for i, record := range records {
		if i == 0 {
			fmt.Println("Headers:", record)
		} else {
			fmt.Printf("Row %d: %v\n", i, record)
		}
	}

	// 11. DIRECTORY OPERATIONS

	fmt.Println("\n--- Directory Operations ---")

	// Create directory
	err = os.MkdirAll("testdir/subdir", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Println("Created testdir/subdir")
	}

	// List directory contents
	entries, _ := os.ReadDir(".")
	fmt.Println("Files in current directory:")
	for _, entry := range entries {
		info, _ := entry.Info()
		typeStr := "FILE"
		if entry.IsDir() {
			typeStr = "DIR "
		}
		fmt.Printf("  %s %s (%d bytes)\n", typeStr, entry.Name(), info.Size())
	}

	// 12. WALKING DIRECTORY TREE

	fmt.Println("\n--- Walking Directory Tree ---")

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			fmt.Println("Found .txt file:", path)
		}
		return nil
	})

	// 13. TEMPORARY FILES

	fmt.Println("\n--- Temporary Files ---")

	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
	} else {
		fmt.Println("Created temp file:", tmpFile.Name())
		tmpFile.WriteString("Temporary data")
		tmpFile.Close()
		// Clean up temp file
		os.Remove(tmpFile.Name())
		fmt.Println("Removed temp file")
	}

	// 14. CLEANUP

	fmt.Println("\n--- Cleanup ---")

	// Remove created files
	os.Remove("example.txt")
	os.Remove("example_copy.txt")
	os.Remove("buffered.txt")
	os.Remove("sessions.json")
	os.Remove("data.csv")
	os.RemoveAll("testdir")
	fmt.Println("Cleaned up test files")
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
