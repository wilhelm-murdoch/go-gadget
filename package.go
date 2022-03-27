package main

type Package struct {
	Name  string  `json:"name"`
	Files []*File `json:"files"`
}
