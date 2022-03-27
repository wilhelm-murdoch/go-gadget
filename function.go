package main

type Function struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
	Body    string `json:"body"`
	Example string `json:"example,omitempty"`
}
