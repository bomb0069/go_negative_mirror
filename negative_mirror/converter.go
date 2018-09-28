package negative_mirror

import "reflect"

type Converter struct {
	original [][]bool
}

func NewConverter() Converter {
	return Converter{}
}

func NewConverterWithOriginal(original [][]bool) Converter {
	return Converter{original}
}

func (c Converter) isNegativeMirrorOf(compareItem [][]bool) bool {
	negativeMirror := c.NegativeMirrorOf(c.original)
	return reflect.DeepEqual(negativeMirror, compareItem)
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
