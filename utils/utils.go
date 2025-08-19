package utils

import (
	"fmt"
	"log"
	"os"
)
func ReadFile(path string) {
	data, err := os.ReadFile(path) 
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println(string(data))
}
