// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package options

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrUnknownOption = fmt.Errorf("unknown option")       // attempt to store an option that isn't on the whitelist
	ErrWrongType     = fmt.Errorf("wrong type for value") // attempt to store an option using the wrong type of data
)

// Options is a simple list of stored data, indexed by name
type Options map[string]interface{}

// ValidOptions is a list of the options that are valid, and their required
// data type
//
// You can safely use any data type you want as the data type. Just remember
// that you need to prefix it with the package name. e.g.
//
//     wl := make(ValidOptions)
//     wl["options"] = "options.Options"
type ValidOptions map[string]string

// OptionsStore is for embedding in your own data structures
type OptionsStore struct {
	ValidOptions ValidOptions // the whitelist of options that can be stored
	Options      Options
}

// NewOptionsStore() will return a standalone OptionsStore for you to use
func NewOptionsStore(whitelist ValidOptions) *OptionsStore {
	retval := &OptionsStore{
		Options:      make(Options),
		ValidOptions: whitelist,
	}

	return retval
}

// SetOption() will store an option for later retrieval
//
// * If the option isn't in the whitelist, returns ErrUnknownOption
// * If the value is the wrong type, returns ErrWrongType
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

// Option() will retrieve an option of any type from the OptionsStore
//
// Once you have retrieved it, you will need to typecast it yourself to
// the original type. This is safe for you to do so, as the whitelist ensures
// that any option stored has the correct data type.
//
//    wl := make(ValidOptions)
//    wl["foo"] = "string"
//
//    o := NewOptionStore(wl)
//    data, ok := o.Option("foo")
//    if !ok {
//        // ... deal with missing data
//    }
//    foo := data.(string)
func (self *OptionsStore) Option(name string) (interface{}, bool) {
	// do we know this option?
	_, ok := self.ValidOptions[name]
	if !ok {
		return nil, false
	}

	// return the data if we have it
	data, ok := self.Options[name]
	return data, ok
}

// OptionAsBool() retrieves an option from the OptionsStore and returns it as
// a boolean value. The second return value indicates whether the option was
// found or not.
func (self *OptionsStore) OptionAsBool(name string) (bool, bool) {
	// do we know this option?
	requiredType, ok := self.ValidOptions[name]
	if !ok {
		return false, false
	}

	// do we have the option?
	data, ok := self.Options[name]
	if !ok {
		return false, false
	}

	// is the option the right type?
	switch requiredType {
	case "bool":
		return data.(bool), true
	case "int":
		if data.(int) == 0 {
			return false, true
		}
		return true, true
	case "string":
		switch strings.ToLower(data.(string)) {
		case "0", "false":
			return false, true
		default:
			return true, true
		}
	default:
		return false, false
	}
}

// OptionAsInt() retrieves an option from the OptionsStore and returns it as
// an int value. The second return value indicates whether the option was found
// or not.
func (self *OptionsStore) OptionAsInt(name string) (int, bool) {
	// do we know this option?
	requiredType, ok := self.ValidOptions[name]
	if !ok {
		return 0, false
	}

	// do we have the data?
	data, ok := self.Options[name]
	if !ok {
		return 0, false
	}

	// is the option the right type?
	switch requiredType {
	case "int":
		return data.(int), true
	case "bool":
		if data.(bool) == false {
			return 0, true
		}
		return 1, true
	case "string":
		retval, err := strconv.Atoi(data.(string))
		if err != nil {
			return 0, false
		}
		return retval, true
	default:
		return 0, false
	}
}

// OptionAsString() retrieves an option from the OptionsStore and returns it
// as a string. The second return value indicates whether the option was found
// or not.
func (self *OptionsStore) OptionAsString(name string) (string, bool) {
	// do we know this option?
	requiredType, ok := self.ValidOptions[name]
	if !ok {
		return "", false
	}

	// do we have the data at all?
	data, ok := self.Options[name]
	if !ok {
		return "", false
	}

	// is the option the right type?
	switch requiredType {
	case "string":
		return data.(string), true
	case "bool":
		if data.(bool) == true {
			return "true", true
		}
		return "false", true
	case "int":
		return strconv.Itoa(data.(int)), true
	default:
		return "", false
	}
}
