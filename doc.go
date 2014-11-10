// Package options is a simple typesafe databag for storing options inside a
// larger data structure of some kind.
//
// Thread Safety
//
// At the moment, there is no built-in thread safety in this library. The
// normal usage pattern is to set options in your main() before starting any
// Goroutines. There should be no actual race conditions when used this way.
//
// If it proves a problem, I'll add some thread safety in a later release.
package options
