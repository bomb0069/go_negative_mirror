package negative_mirror

import (
	"reflect"
	"testing"
)

func Test_empty_array(t *testing.T) {
	original := [][]bool{}
	expected := [][]bool{}

	converter := NewConverter()

	actual := converter.NegativeMirrorOf(original)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}
