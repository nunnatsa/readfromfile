package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func main() {
	if len(os.Args) > 1 {
		for _, fileName := range os.Args[1:] {
			if err := readFromFile(fileName); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}
		return
	}

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

func readFromFile(fileName string) error {
	reader, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("can't open file %s; %w", fileName, err)
	}
	defer reader.Close()

	err = read(reader)
	if err != nil {
		return fmt.Errorf("faile to get data from file %s; %w", fileName, err)
	}

	return nil
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
