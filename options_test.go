// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package options

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestCanInstantiate(t *testing.T) {
	o := NewOptionsStore(nil)
	assert.Equal(t, 0, len(o.ValidOptions))
	assert.Equal(t, 0, len(o.Options))
}

func TestCanInitialise(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["flag"] = "string"
	o := NewOptionsStore(validOptions)
	assert.Equal(t, 1, len(o.ValidOptions))
}

func TestCanSetBoolOption(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["flag"] = "bool"
	o := NewOptionsStore(validOptions)
	o.SetOption("flag", true)
	assert.Equal(t, 1, len(o.Options))
}

func TestCanGetBoolOptions(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["flag"] = "bool"
	o := NewOptionsStore(validOptions)

	expected := true
	o.SetOption("flag", expected)
	actual, ok := o.OptionAsBool("flag")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)
}

func TestCanSetIntOption(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["flag"] = "int"
	o := NewOptionsStore(validOptions)
	o.SetOption("flag", 1)
	assert.Equal(t, 1, len(o.Options))
}

func TestCanGetIntOptions(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["flag"] = "int"
	o := NewOptionsStore(validOptions)

	expected := 99
	o.SetOption("flag", expected)
	actual, ok := o.OptionAsInt("flag")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)
}

func TestCanSetStringOption(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["flag"] = "string"
	o := NewOptionsStore(validOptions)
	o.SetOption("flag", "test value")
	assert.Equal(t, 1, len(o.Options))
}

func TestCanGetStringOptions(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["flag"] = "string"
	o := NewOptionsStore(validOptions)

	expected := "test value"
	o.SetOption("flag", expected)
	actual, ok := o.OptionAsString("flag")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)
}
