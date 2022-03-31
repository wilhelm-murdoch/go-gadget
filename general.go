package gadget

type General struct {
}

func NewGeneral() *General {
	return (&General{}).Parse()
}

func (g *General) Parse() *General {

	return g
}
