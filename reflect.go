package threadMeta

import "unsafe"

// typesByString returns the subslice of typelinks() whose elements have
// the given string representation.
// It may be empty (no known types with that string) or may have
// multiple elements (multiple types with that string).
//
//go:linkname typesByString reflect.typesByString
func typesByString(string) []unsafe.Pointer

// see @runtime.iface
type iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}
