// Copyright (c) 2014-present Stuart Herbert
// Released under a 3-clause BSD license
package options

import (
	"github.com/bmizerany/assert"
	"reflect"
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

func TestCanGetOptions(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["a"] = "string"
	validOptions["b"] = "bool"
	validOptions["c"] = "uint8"
	o := NewOptionsStore(validOptions)

	expectedA := "test value"
	o.SetOption("a", expectedA)
	actualA, ok := o.Option("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, "string", reflect.TypeOf(actualA).String())
	assert.Equal(t, expectedA, actualA.(string))

	expectedB := false
	o.SetOption("b", expectedB)
	actualB, ok := o.Option("b")
	assert.Equal(t, true, ok)
	assert.Equal(t, "bool", reflect.TypeOf(actualB).String())
	assert.Equal(t, expectedB, actualB.(bool))

	expectedC := uint8(100)
	o.SetOption("c", expectedC)
	actualC, ok := o.Option("c")
	assert.Equal(t, true, ok)
	assert.Equal(t, "uint8", reflect.TypeOf(actualC).String())
	assert.Equal(t, expectedC, actualC.(uint8))
}

type SimpleCustomType uint16

func TestCanStoreSimpleCustomType(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["a"] = "options.SimpleCustomType"
	o := NewOptionsStore(validOptions)

	expectedA := SimpleCustomType(1000)
	err := o.SetOption("a", expectedA)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(o.Options))

	actualA, ok := o.Option("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, "options.SimpleCustomType", reflect.TypeOf(actualA).String())
	assert.Equal(t, expectedA, actualA.(SimpleCustomType))
}

type ComplexCustomType struct {
	Foo string
	Bar uint8
}

func TestCanStoreComplexCustomType(t *testing.T) {
	validOptions := make(ValidOptions)
	validOptions["a"] = "options.ComplexCustomType"
	o := NewOptionsStore(validOptions)

	expectedA := ComplexCustomType{
		Foo: "trout",
		Bar: uint8(99),
	}
	err := o.SetOption("a", expectedA)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(o.Options))

	actualA, ok := o.Option("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, "options.ComplexCustomType", reflect.TypeOf(actualA).String())
	assert.Equal(t, expectedA, actualA.(ComplexCustomType))
}

func TestCanRetrieveBoolAsInt(t *testing.T) {
	wl := make(ValidOptions)
	wl["a"] = "bool"
	o := NewOptionsStore(wl)

	stored1 := false
	expected1 := 0
	stored2 := true
	expected2 := 1

	err := o.SetOption("a", stored1)
	assert.Equal(t, nil, err)

	actual1, ok := o.OptionAsInt("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected1, actual1)

	err = o.SetOption("a", stored2)
	assert.Equal(t, nil, err)

	actual2, ok := o.OptionAsInt("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected2, actual2)
}

func TestCanRetrieveBoolAsString(t *testing.T) {
	wl := make(ValidOptions)
	wl["a"] = "bool"
	o := NewOptionsStore(wl)

	stored1 := false
	expected1 := "false"
	stored2 := true
	expected2 := "true"

	err := o.SetOption("a", stored1)
	assert.Equal(t, nil, err)

	actual1, ok := o.OptionAsString("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected1, actual1)

	err = o.SetOption("a", stored2)
	assert.Equal(t, nil, err)

	actual2, ok := o.OptionAsString("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected2, actual2)
}

func TestCanRetrieveIntAsBool(t *testing.T) {
	wl := make(ValidOptions)
	wl["a"] = "int"
	o := NewOptionsStore(wl)

	stored1 := 0
	expected1 := false
	stored2 := 99
	expected2 := true

	err := o.SetOption("a", stored1)
	assert.Equal(t, nil, err)

	actual1, ok := o.OptionAsBool("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected1, actual1)

	err = o.SetOption("a", stored2)
	assert.Equal(t, nil, err)

	actual2, ok := o.OptionAsBool("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected2, actual2)
}

func TestCanRetrieveIntAsString(t *testing.T) {
	wl := make(ValidOptions)
	wl["a"] = "int"
	o := NewOptionsStore(wl)

	stored1 := 0
	expected1 := "0"
	stored2 := 999999
	expected2 := "999999"

	err := o.SetOption("a", stored1)
	assert.Equal(t, nil, err)

	actual1, ok := o.OptionAsString("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected1, actual1)

	err = o.SetOption("a", stored2)
	assert.Equal(t, nil, err)

	actual2, ok := o.OptionAsString("a")
	assert.Equal(t, true, ok)
	assert.Equal(t, expected2, actual2)
}
