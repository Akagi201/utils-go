package errors_test

import (
	"fmt"
	"strings"
	"testing"

	uerrors "github.com/Akagi201/utilgo/errors"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	err := uerrors.From(errors.New("hello"))

	assert.EqualError(t, err, "hello")

	assert.Nil(t, err.Cause())
	assert.Nil(t, errors.Cause(err))

	stacktrace := err.StackTrace()
	assert.NotNil(t, stacktrace)
	assert.NotEmpty(t, stacktrace)

	nothing, ok := uerrors.Get(err, nil)
	assert.False(t, ok)
	assert.Nil(t, nothing)

	nothing, ok = uerrors.Get(err, "nope")
	assert.False(t, ok)
	assert.Nil(t, nothing)
}

func TestCause(t *testing.T) {
	cause := errors.New("cause")
	middle := errors.Wrap(cause, "middle")
	outer := errors.Wrap(middle, "outer")
	err := uerrors.From(outer)

	assert.EqualError(t, err, "outer: middle: cause")

	assert.Equal(t, cause, errors.Cause(err))
}

func TestFormatNoProps(t *testing.T) {
	cause := errors.New("cause")
	err := uerrors.From(cause)

	assert.Equal(t, "cause", fmt.Sprintf("%v", err))

	formatted := fmt.Sprintf("%+v", err)
	prefix := `cause
github.com/Akagi201/utilgo/errors_test.TestFormatNoProps
	` // The rest of the string will have device-specific components (paths, CPU).

	assert.True(t, strings.HasPrefix(formatted, prefix),
		"expected\n%s\nto have prefix\n%s", formatted, prefix)
	assert.Contains(t, formatted, "github.com/Akagi201/utilgo/errors/errors_test.go:")
}

func TestFormatWithProps(t *testing.T) {
	err := uerrors.From(errors.New("cause")).
		WithValue(struct{ Id int }{1}, 2).
		WithValue("key", "value")

	assert.Equal(t, "[key=value,{1}=2] cause", fmt.Sprintf("%v", err))

	formatted := fmt.Sprintf("%+v", err)
	prefix := `[key=value,{Id:1}=2] cause
github.com/Akagi201/utilgo/errors_test.TestFormatWithProps
	` // The rest of the string will have device-specific components (paths, CPU).

	assert.True(t, strings.HasPrefix(formatted, prefix),
		"expected\n%s\nto have prefix\n%s", formatted, prefix)
	assert.Contains(t, formatted, "github.com/Akagi201/utilgo/errors/errors_test.go:")
}

func TestWithValue(t *testing.T) {
	err := uerrors.From(errors.New("hello"))

	withFooBar := err.WithValue("foo", "bar")
	assertHasPropOnSelf(t, withFooBar, "foo", "bar")
	// Get should only check keys, not values.
	assertDoesNotHaveProp(t, withFooBar, "bar")

	// Change the value.
	withFooBaz := withFooBar.WithValue("foo", "baz")
	assertHasPropOnSelf(t, withFooBaz, "foo", "baz")
	// Original error should remain unchanged.
	assertHasPropOnSelf(t, withFooBar, "foo", "bar")

	// Add another value.
	withFooBazKeyValue := withFooBaz.WithValue("key", "value")
	assertHasPropOnSelf(t, withFooBazKeyValue, "key", "value")
	assertHasPropOnSelf(t, withFooBaz, "foo", "baz")
}

func TestGetIntermediateCauseHasProp(t *testing.T) {
	cause := errors.New("root cause")
	middle := uerrors.From(errors.Wrap(cause, "")).
		WithValue("foo", "bar")
	err := errors.Wrap(middle, "")

	assertHasProp(t, err, "foo", "bar")

	// Wrapping with another PropError shouldn't change the results.
	err = uerrors.From(err)
	assertDoesNotHavePropOnSelf(t, err.(uerrors.ContextError), "foo")
	assertHasProp(t, err, "foo", "bar")
}

func TestGetRootCauseHasProp(t *testing.T) {
	cause := uerrors.From(errors.New("root cause")).
		WithValue("foo", "bar")
	middle := errors.Wrap(cause, "")
	err := errors.Wrap(middle, "")

	assertHasProp(t, err, "foo", "bar")

	// Wrapping with another PropError shouldn't change the results.
	err = uerrors.From(err)
	assertDoesNotHavePropOnSelf(t, err.(uerrors.ContextError), "foo")
	assertHasProp(t, err, "foo", "bar")
}

func assertHasPropOnSelf(t *testing.T, err uerrors.ContextError, key, wantVal interface{}) {
	val, ok := err.Get(key)
	assert.True(t, ok)
	assert.Equal(t, wantVal, val)
	assertHasProp(t, err, key, wantVal)
}

// Like assertHasPropOnSelf but calls the package Get to recurse through causes.
func assertHasProp(t *testing.T, err error, key, wantVal interface{}) {
	assert.Equal(t, wantVal, uerrors.GetOpt(err, key))
	val, ok := uerrors.Get(err, key)
	assert.True(t, ok)
	assert.Equal(t, wantVal, val)
}

func assertDoesNotHavePropOnSelf(t *testing.T, err uerrors.ContextError, key interface{}) {
	val, ok := err.Get(key)
	assert.False(t, ok)
	assert.Nil(t, val)
}

// Like assertDoesNotHavePropOnSelf but calls the package Get to recurse through causes.
func assertDoesNotHaveProp(t *testing.T, err error, key interface{}) {
	assert.Nil(t, uerrors.GetOpt(err, key))
	val, ok := uerrors.Get(err, key)
	assert.False(t, ok)
	assert.Nil(t, val)

	if err, ok := err.(uerrors.ContextError); ok {
		assertDoesNotHavePropOnSelf(t, err, key)
	}
}
