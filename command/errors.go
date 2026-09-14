package command

import "fmt"

type NotEnoughArgs struct {
	Expected int
	Got      int
}

func (e NotEnoughArgs) Error() string {
	return fmt.Sprintf("to few arguments, got: %d want: %d", e.Got, e.Expected)
}

type ErrTooManyArgs struct {
	Expected int
	Got      int
}

func (e ErrTooManyArgs) Error() string {
	return fmt.Sprintf("to many arguments, got: %d want: %d", e.Got, e.Expected)
}
