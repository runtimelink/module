package std

// Tag includes a symbol name along with the type of the symbol.
// Format specification is documented in the package documentation.
type Tag string

// Parse returns the symbol, type and (if the
// tag denotes a function) the call signature.
func (tag Tag) Parse() (string, Type, error) {
	return "", Type{}, nil
}

type Type struct {
	Name string // return value if func

	Func bool   // is function?
	Args []Type // arguments (if function)

	Hash bool       // immutablity marker, true if preceded by '#'
	Free rune       // ownership assertion, one of '$', '&', '*', '+' or '-'
	Test Assertions // memory safety assertions
	Fail string     // symbol to lookup on failure (if function)

	Varg int  // index of the %v argument, -1 if none.
	Skip bool // skip this value macro
}

type Assertions struct {
	Capacity bool
	Inverted bool
	Lifetime Assertion
	Overlaps Assertion
	SameType Assertion
	SameSize Assertion
	Equality Assertion
	MoreThan Assertion
	LessThan Assertion
}

type Assertion struct {
	Check bool
	Index int    // of the argument being referred to. if greater than zero ignore const and value.
	Const string // C standard constant (or supported macro) name.
	Value int    // integer value
}
