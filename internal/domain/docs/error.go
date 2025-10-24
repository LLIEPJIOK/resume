package docs

import "fmt"

type ErrInvalidLength struct {
	Expected int
	Actual   int
}

func NewErrInvalidLength(expected, actual int) error {
	return ErrInvalidLength{
		Expected: expected,
		Actual:   actual,
	}
}

func (e ErrInvalidLength) Error() string {
	return fmt.Sprintf("invalid length: expected %d, got %d", e.Expected, e.Actual)
}

type ErrInvalidType struct {
	Type  contentType
	Index int
}

func NewErrInvalidType(tpe contentType, index int) error {
	return ErrInvalidType{
		Type:  tpe,
		Index: index,
	}
}

func (e ErrInvalidType) Error() string {
	return fmt.Sprintf("invalid type %q for index %d", e.Type, e.Index)
}
