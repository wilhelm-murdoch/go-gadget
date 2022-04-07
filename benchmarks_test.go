package gadget_test

import (
	"testing"

	"github.com/wilhelm-murdoch/go-collection"
	"github.com/wilhelm-murdoch/go-gadget"
)

func BenchmarkGadget(b *testing.B) {
	var packages *collection.Collection[*gadget.Package]
	for i := 0; i < b.N; i++ {
		packages = collection.New[*gadget.Package]()

		for _, path := range gadget.WalkGoFiles(".") {
			f, _ := gadget.NewFile(path)

			p := packages.Find(func(i int, p *gadget.Package) bool {
				return p.Name == f.Package
			})
			if p != nil {
				p.Files.Push(f)
				continue
			}

			p = gadget.NewPackage(f.Package)
			p.Files.Push(f)
			packages.Push(p)
		}
	}
}
