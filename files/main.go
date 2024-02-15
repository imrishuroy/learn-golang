package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type ChargeSession struct {
	Id        string `json:"id"`
	Watts     int    `json:"watts"`
	Vin       string `json:"vin"`
	Timestamp string `json:"timestamp"`
}

func main() {

	// reading csv
	csvFP := "session.csv"
	file, err := os.Open(csvFP)
	if err != nil {
		fmt.Println(err)
		return

	}
	defer file.Close()
	reader := csv.NewReader(file)
	session, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return

	}
	for i := 0; i < len(session); i++ {
		for j := 0; j < len(session[i]); j++ {
			fmt.Println(session[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()

	// reading json
	jsonFP := "session.json"
	var cs ChargeSession
	file, err = os.Open(jsonFP)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cs)

	filePath := "test.txt"
	// Create a file
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()

	// Write to a file
	data := []byte("Hey, Everyone I am a writter\nI write code and I write poems")
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File created and written to successfully", filePath)

	// read from a file
	data, err1 := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(string(data))

	// Check if file exists
	_, err3 := os.Stat("test.txt")
	if os.IsNotExist(err3) {
		println("File does not exist")

	} else {
		println("File exists")
	}
}
