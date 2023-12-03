package fileutils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileLines(filename string) ([]string, error) {
	var (
		err    error
		f      *os.File
		result []string
	)

	if f, err = os.Open(filename); err != nil {
		return nil, fmt.Errorf("unable to open file %s: %w", filename, err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}
