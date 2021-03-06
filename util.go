package gadget

import (
	"bufio"
	"go/ast"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
func WalkGoFiles(path string) (files []string) {
	pattern, _ := regexp.Compile(`.+\.go$`) // Suppress errors as this expression will never fail.
	filepath.WalkDir(path, func(path string, dir fs.DirEntry, err error) error {
		if err == nil && pattern.MatchString(dir.Name()) {
			files = append(files, path)
		}
		return nil
	})

	return files
}

// AdjustSource is a convenience function that strips the opening and closing
// braces of a function's ( or other things ) body and removes the first `\t`
// character on each remaining line.
func AdjustSource(source string, adjustBraces bool) string {
	// Remove first leading tab character:
	pattern := regexp.MustCompile(`(?m)^\t{1}`)
	source = pattern.ReplaceAllString(source, "")

	if adjustBraces {
		source = source[:len(source)-1] // Remove trailing } brace
		source = source[1:]             // Remove leading { brace
	}

	return strings.TrimSpace(source) // Trim all leading and trailing space
}

// walker implements a walker interface used to traversing syntax trees.
type Walker func(ast.Node) bool

// Visit steps through each node within the specified syntax tree.
func (w Walker) Visit(node ast.Node) ast.Visitor {
	if w(node) {
		return w
	}
	return nil
}
