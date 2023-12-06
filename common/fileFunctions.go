package common

import (
	"bufio"
	"os"
)

func FileToSlice(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
