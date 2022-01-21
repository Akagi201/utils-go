// Package errors provides a standard error compatible error implementation with key:value context
// It is based on github.com/pkg/errors package: errors returned by this package support both StackTrace() and Cause().
// They also support fmt.Formatter – key:value pairs are formatted using the format flags passed to Format,
// and the wrapped error is asked to format itself.
package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// causer for pkg/errors
type causer interface {
	Cause() error
}

// stackTracer for pkg/errors
type stackTracer interface {
	StackTrace() errors.StackTrace
}

// contexter for key:value context
type contexter interface {
	// Returns the value associated with the given key on the current error.
	// Should *NOT* recurse into the error's cause, if it has one.
	Get(key any) (value any, ok bool)
}

// ContextError An implementation of the standard error interface with key:value context.
// ContextError is immutable.
//
// This interface should never be used as a return *type* for a function.
//
// Good:
//		func do() error { return errors.From(…) }
// Bad:
//		func do() ContextError { return errors.From(…) }
//
type ContextError interface {
	error
	causer
	stackTracer
	contexter
	fmt.Formatter

	// Returns a copy of this ContextError with the specified key/value pair.
	// Do not modify the current ContextError.
	WithValue(key, value any) ContextError

	// If the wrapped error implements fmt.Formatter, this method should delegate directly
	// to it.
	formatBaseError(f fmt.State, c rune)
}

// From teturns a ContextError that can be used to set key:value context on err.
// The original err is not modified.
//
// Intended to be used in a fluent style, like:
// 		return From(err).
//			WithValue("someKey", someValue).
// 			WithValue("anotherKey", anotherValue)
func From(err error) ContextError {
	return baseError{err}
}

// Get returns the value associated with key using the following rules to resolve key:
//   - If the key exists on the current error, return the value set by the last call to WithValue.
//   - Else, if the current error has a cause, recursively look for the key on the cause.
func Get(err error, key any) (any, bool) {
	if err == nil {
		return nil, false
	}

	if err, ok := err.(contexter); ok {
		if val, ok := err.Get(key); ok {
			return val, true
		}
	}

	if err, ok := err.(causer); ok {
		return Get(err.Cause(), key)
	}

	return nil, false
}

// GetOpt is the same as Get but just returns nil if the key isn't set, to make it
// more convenient to use in single-value contexts.
func GetOpt(err error, key any) any {
	if val, ok := Get(err, key); ok {
		return val
	}
	return nil
}

// baseError implementation of ContextError that just delegates most calls to the underlying error
// if it supports them.
type baseError struct {
	error
}

func (e baseError) Cause() error {
	if e, ok := e.error.(causer); ok {
		return e.Cause()
	}
	return nil
}

func (e baseError) StackTrace() errors.StackTrace {
	if e, ok := e.error.(stackTracer); ok {
		return e.StackTrace()
	}
	return nil
}

func (e baseError) Get(key any) (any, bool) {
	if e, ok := e.error.(contexter); ok {
		return e.Get(key)
	}
	return nil, false
}

func (e baseError) WithValue(key, value any) ContextError {
	return &kvError{e, key, value}
}

func (e baseError) Format(f fmt.State, c rune) {
	e.formatBaseError(f, c)
}

func (e baseError) formatBaseError(f fmt.State, c rune) {
	if e, ok := e.error.(fmt.Formatter); ok {
		e.Format(f, c)
		return
	}
	fmt.Fprint(f, e.error)
}

// kvError a wrapper around an ContextError that associates a key with a value.
type kvError struct {
	ContextError
	key, value any
}

func (e *kvError) Get(key any) (any, bool) {
	if key == e.key {
		return e.value, true
	}
	return e.ContextError.Get(key)
}

func (e *kvError) WithValue(key, value any) ContextError {
	// Override here so the wrapped error is this object, not the embedded ContextError.
	return &kvError{e, key, value}
}

func (e *kvError) formatBaseError(f fmt.State, c rune) {
	e.ContextError.formatBaseError(f, c)
}

func (e *kvError) Format(f fmt.State, c rune) {
	fmt.Fprint(f, "[")
	e.formatInner(f, c)
	fmt.Fprint(f, "] ")
	e.formatBaseError(f, c)
}

func (e *kvError) formatInner(f fmt.State, c rune) {
	format := "%"

	if f.Flag('+') {
		format += "+"
	} else if f.Flag('#') {
		format += "#"
	}
	format = string(append([]rune(format), c))

	fmt.Fprintf(f, format+"="+format, e.key, e.value)

	if err, ok := e.ContextError.(*kvError); ok {
		fmt.Fprint(f, ",")
		err.formatInner(f, c)
	}
}
