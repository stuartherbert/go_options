// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package options

import (
	"fmt"
	"reflect"
)

var (
	ErrUnknownOption = fmt.Errorf("unknown option")
	ErrWrongType     = fmt.Errorf("wrong type for value")
)

// Options is
type Options map[string]interface{}

// ValidOptions is a list of the options that are valid, and their required
// data type
type ValidOptions map[string]string

type OptionsStore struct {
	ValidOptions ValidOptions
	Options      Options
}

func NewOptionsStore(validOptions ValidOptions) *OptionsStore {
	retval := &OptionsStore{
		Options:      make(Options),
		ValidOptions: validOptions,
	}

	return retval
}

func (self *OptionsStore) SetOption(name string, value interface{}) error {
	// is this a valid option?
	requiredType, ok := self.ValidOptions[name]
	if !ok {
		return ErrUnknownOption
	}

	// do we have the right data type?
	actualType := reflect.TypeOf(value).String()
	if requiredType != actualType {
		return ErrWrongType
	}

	// at this point, it is safe to store the option
	self.Options[name] = value
	return nil
}

func (self *OptionsStore) GetOptionString(name string) (string, bool) {
	// do we know this option?
	requiredType, ok := self.ValidOptions[name]
	if !ok {
		return "", false
	}

	if requiredType != "string" {
		return "", false
	}

	return self.Options[name].(string), true
}