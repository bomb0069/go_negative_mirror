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

func Test_only_one_member_it_mirror(t *testing.T) {
	original := [][]bool{{true}}
	expected := [][]bool{{false}}

	converter := NewConverter()

	actual := converter.NegativeMirrorOf(original)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

func Test_two_member_one_row_it_mirror(t *testing.T) {
	original := [][]bool{{true, false}}
	expected := [][]bool{{true, false}}

	converter := NewConverter()

	actual := converter.NegativeMirrorOf(original)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}
