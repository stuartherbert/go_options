// Package options is a simple typesafe databag for storing options inside a
// larger data structure of some kind.
//
// Whitelist
//
// At the heart of the OptionsStore is a user-supplied whitelist of which
// options are allowed, and their data types:
//
//    wl := make(ValidOptions)
//    wl["tom"] = "string"
//    wl["dick"] = "int"
//    wl["harry"] = "bool"
//    wl["kate"] = "option.Options"
//
//    os := NewOptionsStore(wl)
//
// You can use any Golang datatype. If you use data types defined in packages,
// just remember to prefix the datatype with the package name!
//
// Storing / Retrieving Options
//
// The OptionsStore lets you store any option that validates against the
// whitelist:
//
//    os := NewOptionsStore(wl)
//    os.SetOption("tom", "hardy")
//    os.SetOption("dick", 99)
//    os.SetOption("harry", true)
//
// There's a generic way to get any option stored in the OptionsStore:
//
//    option, ok := os.Option("tom")
//    if !ok {
//        // ... deal with missing option
//    }
//    data := option.(string)
//
// You can safely typecast the retrieved data. It can only be the type that
// matches the whitelist entry.
//
// There are also some helpers to save you having to do the typecasting:
//
//    data1 := os.OptionAsString("tom")
//    data2 := os.OptionAsInt("dick")
//    data3 := os.OptionAsBool("harry")
//
// Converting Options
//
// You can use the helpers to typecast the stored option into a different
// type (within reason).
//
//    // os.OptionAsBool() will return false for any of these
//    os.SetOption("tom", "false")
//    os.SetOption("tom", "fAlSe")
//    os.SetOption("dick", 0)
//
//    // os.OptionAsInt() will return 0 for this
//    os.SetOption("tom", "0")
//    os.SetOption("harry", false)
//
//    // os.OptionAsInt() will return 1 for this
//    os.SetOption("tom", "1")
//    os.SetOption("harry", true)
//
//    // os.OptionAsInt() will return strconv.Atoi() for this
//    os.SetOption("tom", "999999")
//
//    // os.OptionAsString() will return "false" for this
//    os.SetOption("harry", false)
//
//    // os.OptionAsString() will return "true" for this
//    os.SetOption("harry", true)
//
//    // os.OptionAsString() will return strconv.Itoa() for this
//    os.SetOption("dick", 999999)
//
// Thread Safety
//
// At the moment, there is no built-in thread safety in this library. The
// normal usage pattern is to set options in your main() before starting any
// Goroutines. There should be no actual race conditions when used this way.
//
// If it proves a problem, I'll add some thread safety in a later release.
package options
