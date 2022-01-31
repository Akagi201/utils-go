package flags_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/Akagi201/utils-go/flags"
)

var (
	tags       flags.Array
	attributes flags.Map
)

func init() {
	flag.Var(&tags, "tag", "Tag to add.")
	flag.Var(&attributes, "attribute", "Attribute to add.")
}

func Example() {
	os.Args = []string{"-tag=1", "-tag=2", "-attribute", "foo:bar", "-attribute", "baz:boo"}
	flag.Parse()

	fmt.Println("tags provided:", tags)
	fmt.Println("attributes provided:", attributes)

	// When started with:
	// -tag=1 -tag=2 -attribute foo:bar -attribute baz:boo
	// Outputs:
	// tags provided: [1 2]
	// attributes provided: map[foo:bar baz:boo]
}
