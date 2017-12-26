// Package fileutil provides functions for doing things with files, like reading them
// line by line, etc
package files

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var propertySplittingRegex = regexp.MustCompile(`\s*=\s*`)

// ReadLinesChannel reads a text file line by line into a channel.
//
//   c, err := fileutil.ReadLinesChannel(fileName)
//   if err != nil {
//      log.Fatalf("readLines: %s\n", err)
//   }
//   for line := range c {
//      fmt.Printf("  Line: %s\n", line)
//   }
//
// nil is returned (with the error) if there is an error opening the file
//
func ReadLinesChannel(filePath string) (<-chan string, error) {
	c := make(chan string)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	go func() {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()
	return c, nil
}

// ReadLinesSlice reads a text file line by line into a slice of strings.
// Not recommended for use with very large files due to the memory needed.
//
//   lines, err := fileutil.ReadLinesSlice(filePath)
//   if err != nil {
//       log.Fatalf("readLines: %s\n", err)
//   }
//   for i, line := range lines {
//       fmt.Printf("  Line: %d %s\n", i, line)
//   }
//
// nil is returned if there is an error opening the file
//
func ReadLinesSlice(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteLinesSlice writes the given slice of lines to the given file.
func WriteLinesSlice(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// Exists returns whether or not the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// TempFileName generates a temporary filename for use in testing or whatever
func TempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
}

// Reads name-value pairs from a properties file
// Property file has lines in the format of "name = value" (leading and trailing spaces are ignored)
func ReadPropertiesFile(fileName string) (map[string]string, error) {
	c, err := ReadLinesChannel(fileName)
	if err != nil {
		return nil, err
	}

	properties := make(map[string]string)
	for line := range c {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			// Ignore this line
		} else if len(line) == 0 {
			// Ignore this line
		} else {
			parts := propertySplittingRegex.Split(line, 2)
			properties[parts[0]] = parts[1]
		}
	}

	return properties, nil
}

// Pwd returns the present working directory
func Pwd() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
