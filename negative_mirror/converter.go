package negative_mirror

type Converter struct {
}

func NewConverter() Converter {
	return Converter{}
}

func (c Converter) NegativeMirrorOf(original [][]bool) [][]bool {
	negativeMirror := make([][]bool, len(original))
	for outerIndex, innerArray := range original {
		innerSize := len(innerArray)
		nagativeMirrorRow := make([]bool, innerSize)
		for innerIndex, value := range innerArray {
			nagativeMirrorRow[(innerSize-1)-innerIndex] = !value
		}
		negativeMirror[outerIndex] = nagativeMirrorRow
	}
	return negativeMirror
}
