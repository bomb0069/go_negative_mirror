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

func Test_three_member_one_row_it_mirror(t *testing.T) {
	original := [][]bool{{true, true, false}}
	expected := [][]bool{{true, false, false}}

	converter := NewConverter()

	actual := converter.NegativeMirrorOf(original)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

func Test_false_true_false_one_row_it_mirror(t *testing.T) {
	original := [][]bool{{false, true, false}}
	expected := [][]bool{{true, false, true}}

	converter := NewConverter()

	actual := converter.NegativeMirrorOf(original)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

func Test_original_problem_in_first_row_it_mirror(t *testing.T) {
	original := [][]bool{{false, true, true, false, true}}
	expected := [][]bool{{false, true, false, false, true}}

	converter := NewConverter()

	actual := converter.NegativeMirrorOf(original)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

func Test_original_problem_it_mirror(t *testing.T) {
	original := [][]bool{
		{false, true, true, false, true},
		{true, true, false, false, true},
		{false, false, true, true, true},
	}
	expected := [][]bool{
		{false, true, false, false, true},
		{false, true, true, false, false},
		{false, false, false, true, true},
	}

	converter := NewConverter()

	actual := converter.NegativeMirrorOf(original)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%v != %v", expected, actual)
	}
}

func Test_original_problem_is_negative_mirror(t *testing.T) {
	original := [][]bool{
		{false, true, true, false, true},
		{true, true, false, false, true},
		{false, false, true, true, true},
	}
	negative_mirror := [][]bool{
		{false, true, false, false, true},
		{false, true, true, false, false},
		{false, false, false, true, true},
	}

	orginalConverter := NewConverterWithOriginal(original)

	if !orginalConverter.isNegativeMirrorOf(negative_mirror) {
		t.Errorf("%v is not negative mirror of %v", negative_mirror, original)
	}
}
