package negative_mirror

type Converter struct {
}

func NewConverter() Converter {
	return Converter{}
}

func (c Converter) NegativeMirrorOf(original [][]bool) [][]bool {
	return original
}
