package errors_test

import (
	"fmt"
	"io"
	"time"

	uerrors "github.com/Akagi201/utilgo/errors"
	"github.com/pkg/errors"
)

func Example() {
	var err error = errors.New("fail")
	var errWithProp error = uerrors.From(err).
		WithValue("key", "value").
		WithValue("foo", "bar")

	fmt.Println("key:", uerrors.GetOpt(errWithProp, "key"))
	fmt.Println("foo:", uerrors.GetOpt(errWithProp, "foo"))

	// Output:
	// key: value
	// foo: bar
}

func Example_cause() {
	var rootCause error = uerrors.From(errors.New("root cause")).
		WithValue("rootKey", "rootValue")

	var wrapped error = uerrors.From(errors.Wrap(rootCause, "wrapped")).
		WithValue("wrappedKey", "wrappedValue")

	fmt.Println("rootKey:", uerrors.GetOpt(wrapped, "rootKey"))
	fmt.Println("wrappedKey:", uerrors.GetOpt(wrapped, "wrappedKey"))

	// Output:
	// rootKey: rootValue
	// wrappedKey: wrappedValue
}

func Example_overriding() {
	var rootCause error = uerrors.From(errors.New("root cause")).
		WithValue("key", "rootValue")

	var wrapped error = uerrors.From(errors.Wrap(rootCause, "wrapped")).
		WithValue("key", "wrappedValue")

	fmt.Println("key:", uerrors.GetOpt(rootCause, "key"))
	fmt.Println("key:", uerrors.GetOpt(wrapped, "key"))

	// Output:
	// key: rootValue
	// key: wrappedValue
}

func Example_format() {
	type state struct {
		Url      string
		Duration time.Duration
	}

	err := uerrors.From(io.EOF).
		WithValue("filename", "/tmp/stuff").
		WithValue("state", state{
			"https://example.com",
			2 * time.Second,
		})

	fmt.Printf("%v\n", err)
	fmt.Printf("%+v\n", err)
	fmt.Printf("%#v\n", err)

	// Output:
	// [state={https://example.com 2s},filename=/tmp/stuff] EOF
	// [state={Url:https://example.com Duration:2s},filename=/tmp/stuff] EOF
	// ["state"=errors_test.state{Url:"https://example.com", Duration:2000000000},"filename"="/tmp/stuff"] EOF
}
