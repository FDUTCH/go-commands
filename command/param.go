package command

import (
	"fmt"
	"slices"
	"strconv"
)

// ParamType represents param type.
type ParamType byte

const (
	ParamTypeBool = ParamType(iota)
	ParamTypeInt
	ParamTypeFloat
	ParamTypeString
	ParamTypeEnum
)

// ParamDescriptor contains information that describes command parameter.
type ParamDescriptor struct {
	Name     string
	Type     ParamType
	Optional bool
	EnumDesc string

	val              any // pointer the primitive type
	hasVal           *bool
	availableOptions []string
}

// AvailableEnumOptions returns available enum options of enum type.
func (p ParamDescriptor) AvailableEnumOptions() []string {
	return slices.Clone(p.availableOptions)
}

func (p ParamDescriptor) parse(val string) error {
	if p.Optional {
		*p.hasVal = true
	}
	switch p.Type {
	case ParamTypeBool:
		return p.parseBool(val)
	case ParamTypeInt:
		return p.parseInt(val)
	case ParamTypeFloat:
		return p.parseFloat(val)
	case ParamTypeString:
		return p.parseString(val)
	case ParamTypeEnum:
		return p.parseEnum(val)
	default:
		return fmt.Errorf("unknown param type: %v", p.Type)
	}
}

func (p ParamDescriptor) parseBool(val string) error {
	var err error
	ptr := p.val.(*bool)
	*ptr, err = strconv.ParseBool(val)
	return err
}

func (p ParamDescriptor) parseInt(val string) error {
	var err error
	ptr := p.val.(*int)
	*ptr, err = strconv.Atoi(val)
	return err
}

func (p ParamDescriptor) parseFloat(val string) error {
	var err error
	ptr := p.val.(*float64)
	*ptr, err = strconv.ParseFloat(val, 64)
	return err
}

func (p ParamDescriptor) parseString(val string) error {
	*(p.val.(*string)) = val
	return nil
}

func (p ParamDescriptor) parseEnum(val string) error {
	if !slices.Contains(p.availableOptions, val) {
		return fmt.Errorf("%s is not a valid option for: %s", val, p.EnumDesc)
	}
	*(p.val.(*string)) = val
	return nil
}
