package gadget_test

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget"
)

func TestGetLinesFromFile(t *testing.T) {
	var expected string
	var found []byte

	expected = "package main"
	found = gadget.GetLinesFromFile("cmd/main.go", 1, 1)
	assert.Equal(t, expected, strings.TrimSpace(string(found)), "Expected the first line of target file to equal `%s`, but got `%s` instead.", expected, found)

	expected = "package gadget_test"
	found = gadget.GetLinesFromFile("util_test.go", 1, 1)
	assert.Equal(t, expected, strings.TrimSpace(string(found)), "Expected the first line of target file to equal `%s`, but got `%s` instead.", expected, found)

	found = gadget.GetLinesFromFile("not_found.go", 1, 1)
	assert.Equal(t, []byte{}, found, "Expected to return nothing as the file does not exist, but got `%s` instead.", found)

	found = gadget.GetLinesFromFile("util.go", 9999, 9999)
	assert.Equal(t, []byte{}, found, "Expected to return nothing as the range does not exist, but got `%s` instead.", found)

	found = gadget.GetLinesFromFile("util.go", 1, 50)
	lines := bytes.Count(found, []byte{'\n'})
	assert.Equal(t, lines, 50, "Expected to return %d lines, but got %d instead.", 50, lines)
}

func TestWalkGoFiles(t *testing.T) {
	var files []string
	files = gadget.WalkGoFiles("cmd/")
	assert.Equal(t, len(files), 2, "Expected to find 2 files, but got %d instead.", len(files))

	files = gadget.WalkGoFiles("cmd/not_found/")
	assert.Equal(t, len(files), 0, "Expected to find 0 files, but got %d instead.", len(files))

	files = gadget.WalkGoFiles("cmd/main.go")
	assert.Equal(t, len(files), 1, "Expected to find 1 files, but got %d instead.", len(files))
	assert.Contains(t, files, "cmd/main.go", "Expected to find `cmd/main.go`, but found `%s` instead.", files[0])
}

func TestAdjustSource(t *testing.T) {
	var adjusted string
	source := `{
	this is a line
	this is another line
}`

	sourceBracesKeep := `{
this is a line
this is another line
}`

	sourceBracesDrop := `this is a line
this is another line`

	adjusted = gadget.AdjustSource(source, false)
	assert.Equal(t, adjusted, sourceBracesKeep, "Expected to keep opening and closing braces while removing initial \\t on each line")

	adjusted = gadget.AdjustSource(source, true)
	assert.Equal(t, adjusted, sourceBracesDrop, "Expected to drop opening and closing braces while removing initial \\t on each line")
}

func TestWalker_Visit(t *testing.T) {
	tokenSet := token.NewFileSet()
	astFile, _ := parser.ParseFile(tokenSet, "noop.go", "", 0)

	(&TestWalker{astFile}).walk(func(node ast.Node) bool {
		fmt.Println(node)
		return true
	})
}

type TestWalker struct {
	astFile *ast.File
}

func (tw *TestWalker) walk(fn func(ast.Node) bool) {
	ast.Walk(gadget.Walker(fn), tw.astFile)
}
