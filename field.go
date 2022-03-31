package gadget

type Field struct {
	Name       string `json:"name"`
	IsExported bool   `json:"is_exported"`
	Line       int    `json:"line"`
	Body       string `json:"body"`
}

func (f *Field) String() string {
	return f.Name
}
