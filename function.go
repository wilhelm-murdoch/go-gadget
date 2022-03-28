package main

type Function struct {
	Package  string `json:"package"`
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Comment  string `json:"comment,omitempty"`
	Body     string `json:"body"`
	Example  string `json:"example,omitempty"`
	Exported bool   `json:"exported"`
}

func NewFunction(name, pkg, fileName, filePath string) *Function {
	return &Function{
		Package:  pkg,
		Name:     name,
		FilePath: filePath,
		FileName: fileName,
	}
}

func (f *Function) formatSource(source string) string {
	return source
}

func (f *Function) SetExample(source string) {
	f.Example = f.formatSource(source)
}

func (f *Function) SetBody(source string) {
	f.Body = f.formatSource(source)
}
