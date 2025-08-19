package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)
/*
Parse the labels and paths into a map to keep track of available shortcuts
*/
func ParseLabelsPaths(filePath string) (map[string]string) {
	var labelsPaths = make(map[string]string)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	
	// Parse labels and paths from file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parsed := strings.SplitN(line, ":", 2)
		labelsPaths[parsed[0]] = parsed[1]
	}
	return labelsPaths
}
