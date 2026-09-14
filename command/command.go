package command

import "slices"

// Command instance should be used in per user scope -
// only one user should have access to the current instance at the same time.
// Command should never run from 2 (or more) different goroutines.
type Command struct {
	sub     *Command
	subName string

	descriptors []ParamDescriptor
	run         func()
}

// Descriptors returns slice of param descriptors of the current Command.
func (c *Command) Descriptors() []ParamDescriptor {
	return slices.Clone(c.descriptors)
}

// Run parses args into a values that can be accessed with Result objects obtained when command is build,
// and runs "run" func of the Command.
func (c *Command) Run(args []string) (err error, place int) {
	if len(args) == 0 {
		if len(c.descriptors) == 0 {
			c.run()
			return nil, 0
		}
		return NotEnoughArgs{Expected: len(c.descriptors)}, 0
	}

	if args[0] == c.subName {
		return c.sub.Run(args[1:])
	}

	if len(args) > len(c.descriptors) {
		return ErrTooManyArgs{Got: len(args), Expected: len(c.descriptors)}, 0
	}

	for idx, desc := range c.descriptors {
		if idx >= len(args) {
			if desc.Optional {
				*desc.hasVal = false
				continue
			} else {
				return NotEnoughArgs{len(c.descriptors), len(args)}, idx
			}
		}

		if err := c.descriptors[idx].parse(args[idx]); err != nil {
			return err, idx
		}
	}

	c.run()
	return nil, 0
}
