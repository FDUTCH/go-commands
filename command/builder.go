package command

// Builder builds command.
type Builder struct {
	sub          *Command
	subName      string
	descriptors  []ParamDescriptor
	optionalMode bool
	aliases      []string
}

// AddSubCommand adds subcommand.
func (b *Builder) AddSubCommand(cmd *Command, name string) {
	b.sub = cmd
	b.subName = name
}

// AddAlias adds command alias.
func (b *Builder) AddAlias(alias string) {
	b.aliases = append(b.aliases, alias)
}

// EnableOptionalMode enables optional mode.
// After enabling, all future params will be optional.
func (b *Builder) EnableOptionalMode() {
	b.optionalMode = true
}

// Bool adds bool param.
func (b *Builder) Bool(name string) Result[bool] {
	return b.addDescriptor[bool](ParamDescriptor{
		Name: name,
		Type: ParamTypeBool,
	})
}

// Int adds int param.
func (b *Builder) Int(name string) Result[int] {
	return b.addDescriptor[int](ParamDescriptor{
		Name: name,
		Type: ParamTypeInt,
	})
}

// Float adds float param.
func (b *Builder) Float(name string) Result[float64] {
	return b.addDescriptor[float64](ParamDescriptor{
		Name: name,
		Type: ParamTypeFloat,
	})
}

// String adds string param.
func (b *Builder) String(name string) Result[string] {
	return b.addDescriptor[string](ParamDescriptor{
		Name: name,
		Type: ParamTypeString,
	})
}

// Enum adds enum param.
func (b *Builder) Enum(name, desc string, options ...string) Result[string] {
	return b.addDescriptor[string](ParamDescriptor{
		Name:             name,
		Type:             ParamTypeEnum,
		availableOptions: options,
		EnumDesc:         desc,
	})
}

func (b *Builder) addDescriptor[T any](descriptor ParamDescriptor) Result[T] {
	if b.optionalMode {
		descriptor.Optional = true
		descriptor.hasVal = new(bool)
	}
	ptr := new(T)
	descriptor.val = ptr
	b.descriptors = append(b.descriptors, descriptor)
	return Result[T]{ptr, descriptor.hasVal}
}

// Build builds Command.
func (b *Builder) Build(description string, run func()) *Command {
	return &Command{
		description: description,
		aliases:     b.aliases,
		sub:         b.sub,
		subName:     b.subName,
		descriptors: b.descriptors,
		run:         run,
	}
}
