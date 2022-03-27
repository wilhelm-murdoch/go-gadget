package main

type File struct {
	Path      string      `json:"path"`
	Name      string      `json:"name"`
	Functions []*Function `json:"functions"`
}
