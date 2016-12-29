// Package htmls contains helper functions simplify working with golang.org/x/net/html package
package htmls

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// MatchFunc matches HTML nodes, should return true when a desired node is found.
type MatchFunc func(*html.Node) bool

// MatchAtom returns a MatchFunc that matches a Node with the specified Atom.
func MatchAtom(a atom.Atom) MatchFunc {
	return func(n *html.Node) bool {
		return n.DataAtom == a
	}
}

// FindAll returns all nodes which match the provided MatchFunc. After discovering a matching
// node, it will _not_ discover matching subnodes of that node.
func FindAll(node *html.Node, mf MatchFunc) []*html.Node {
	return findAllInternal(node, mf, false)
}

// FindAllNested returns all nodes which match the provided MatchFunc and _will_ discover
// matching subnodes of matching nodes.
func FindAllNested(node *html.Node, mf MatchFunc) []*html.Node {
	return findAllInternal(node, mf, true)
}

// Find returns the first node which matches the MatchFunc using depth-first search.
// If no node is found, ok will be false.
//
//     root, err := html.Parse(resp.Body)
//     if err != nil {
//         // handle error
//     }
//     mf := func(n *html.Node) bool {
//         return n.DataAtom == atom.Body
//     }
//     body, ok := scrape.Find(root, mf)
func Find(node *html.Node, mf MatchFunc) (n *html.Node, ok bool) {
	if mf(node) {
		return node, true
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		n, ok := Find(c, mf)
		if ok {
			return n, true
		}
	}
	return nil, false
}

// FindParent searches up HTML tree from the current node until either a
// match is found or the top is hit.
func FindParent(node *html.Node, mf MatchFunc) (n *html.Node, ok bool) {
	for p := node.Parent; p != nil; p = p.Parent {
		if mf(p) {
			return p, true
		}
	}
	return nil, false
}

// Text returns text from all descendant text nodes joined.
// For control over the join function, see TextJoin.
func Text(node *html.Node) string {
	joiner := func(s []string) string {
		n := 0
		for i := range s {
			trimmed := strings.TrimSpace(s[i])
			if trimmed != "" {
				s[n] = trimmed
				n++
			}
		}
		return strings.Join(s[:n], " ")
	}
	return TextJoin(node, joiner)
}

// TextJoin returns a string from all descendant text nodes joined by a
// caller provided join function.
func TextJoin(node *html.Node, join func([]string) string) string {
	nodes := FindAll(node, func(n *html.Node) bool { return n.Type == html.TextNode })
	parts := make([]string, len(nodes))
	for i, n := range nodes {
		parts[i] = n.Data
	}
	return join(parts)
}

// Attr returns the value of an HTML attribute.
func Attr(node *html.Node, key string) string {
	for _, a := range node.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

// ByTag returns a MatchFunc which matches all nodes of the provided tag type.
//
//     root, err := html.Parse(resp.Body)
//     if err != nil {
//         // handle error
//     }
//     title, ok := scrape.Find(root, scrape.ByTag(atom.Title))
func ByTag(a atom.Atom) MatchFunc {
	return func(node *html.Node) bool { return node.DataAtom == a }
}

// ByID returns a MatchFunc which matches all nodes with the provided id.
func ByID(id string) MatchFunc {
	return func(node *html.Node) bool { return Attr(node, "id") == id }
}

// ByClass returns a MatchFunc which matches all nodes with the provided class.
func ByClass(class string) MatchFunc {
	return func(node *html.Node) bool {
		classes := strings.Fields(Attr(node, "class"))
		for _, c := range classes {
			if c == class {
				return true
			}
		}
		return false
	}
}

// findAllInternal encapsulates the node tree traversal
func findAllInternal(node *html.Node, mf MatchFunc, searchNested bool) []*html.Node {
	matched := []*html.Node{}

	if mf(node) {
		matched = append(matched, node)

		if !searchNested {
			return matched
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		found := findAllInternal(c, mf, searchNested)
		if len(found) > 0 {
			matched = append(matched, found...)
		}
	}
	return matched
}

// FindNextSibling returns the first node which matches the MatchFunc using next sibling search.
// If no node is found, ok will be false.
//
//     root, err := html.Parse(resp.Body)
//     if err != nil {
//         // handle error
//     }
//     mf := func(n *html.Node) bool {
//         return n.DataAtom == atom.Body
//     }
//     body, ok := scrape.FindNextSibling(root, mf)
func FindNextSibling(node *html.Node, mf MatchFunc) (n *html.Node, ok bool) {

	for s := node.NextSibling; s != nil; s = s.NextSibling {
		if mf(s) {
			return s, true
		}
	}
	return nil, false
}

// FindPrevSibling returns the first node which matches the MatchFunc using previous sibling search.
// If no node is found, ok will be false.
//
//     root, err := html.Parse(resp.Body)
//     if err != nil {
//         // handle error
//     }
//     mf := func(n *html.Node) bool {
//         return n.DataAtom == atom.Body
//     }
//     body, ok := scrape.FindPrevSibling(root, mf)
func FindPrevSibling(node *html.Node, mf MatchFunc) (n *html.Node, ok bool) {
	for s := node.PrevSibling; s != nil; s = s.PrevSibling {
		if mf(s) {
			return s, true
		}
	}
	return nil, false
}
