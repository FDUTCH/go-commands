package command

// Result is a result of command execution.
// It may be optional, and if so "Optional" methods should be used...
type Result[T any] struct {
	val *T
	has *bool
}

// Val returns value.
func (r Result[T]) Val() T {
	if r.has != nil {
		panic("calling Val on optional result")
	}
	return *r.val
}

// Optional returns true if Result is optional.
func (r Result[T]) Optional() bool {
	return r.has != nil
}

// OptionalVal returns Value and true if value is set.
func (r Result[T]) OptionalVal() (T, bool) {
	if r.has == nil {
		panic("calling optional method on non optional result")
	}
	return *r.val, *r.has
}

// OptionalValOr returns value or 'or' value.
func (r Result[T]) OptionalValOr(or T) T {
	if r.has == nil {
		panic("calling optional method on non optional result")
	}
	if *r.has {
		return *r.val
	}
	return or
}
