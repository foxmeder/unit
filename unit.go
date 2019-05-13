package unit

// Unit represents a unit
type Unit float64

// Uniter is a type that can be converted to a Unit.
type Uniter interface {
	Unit() Unit
}
