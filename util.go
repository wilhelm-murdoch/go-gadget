package gadget

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

// GetLinesFromFile creates a byte reader for the file at the target path and
// returns a slice of bytes representing the file content. This slice is
// restricted the lines specified by the `from` and `to` arguments inclusively.
// This will return an empty byte if an empty file, or any error, is encountered.
func GetLinesFromFile(path string, from, to int) []byte {
	var out []byte
	line := 1

	file, err := os.Open(path)
	if err != nil {
		return []byte("")
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		read, err := r.ReadBytes('\n')
		if err == io.EOF && line < from {
			return []byte("")
		}

		if line >= from && line <= to {
			out = append(out, read[:]...)
		} else if from <= line && to <= line {
			break
		}

		line++
	}

	return out
}

// WalkGoFiles recursively moves through the directory tree specified by `path`
// providing a slice of files matching the `*.go` extention. Explicitly
// specifying a file will return that file.
func WalkGoFiles(path *string) (files []string) {
	pattern, err := regexp.Compile(".+\\.go$")
	if err != nil {
		return
	}

	filepath.WalkDir(*path, func(path string, dir fs.DirEntry, err error) error {
		if err == nil && pattern.MatchString(dir.Name()) {
			files = append(files, path)
		}
		return nil
	})

	return
}
