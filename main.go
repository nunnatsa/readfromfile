package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

//var fileName string
//
//func init() {
//	flag.StringVar(&fileName, "input", "", "input file path; if missing, reading from stdin")
//	flag.Parse()
//}

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func main() {
	if len(os.Args) > 1 {
		for _, fileName := range os.Args[1:] {
			readFromFile(fileName)
		}
		return
	}

	//if len(fileName) > 0 {
	//	reader, err = os.Open(fileName)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "can't open file %s; %v\n", fileName, err)
	//		os.Exit(1)
	//	}
	//}

	var (
		reader io.Reader = os.Stdin
		err    error
	)

	err = read(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "faile to get data; %v\n", err)
		os.Exit(1)
	}
}

func readFromFile(fileName string) {
	reader, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s; %v\n", fileName, err)
		os.Exit(1)
	}
	defer reader.Close()
	err = read(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "faile to get data from file %s; %v\n", fileName, err)
		os.Exit(1)
	}
}

func read(reader io.Reader) error {
	users := make([]User, 0)

	dec := json.NewDecoder(reader)

	err := dec.Decode(&users)
	if err != nil {
		return fmt.Errorf("failed to parse the file; %w", err)
	}

	for _, user := range users {
		fmt.Printf("Name: %s\nemail: %s\n", user.Name, user.Email)
	}

	return nil
}
