package htmls_test

import (
	"strings"
	"testing"

	"github.com/Akagi201/utils-go/htmls"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const testHTML = `
<html>
  <body>
    <div class="bigbird">
      <div class="container">
        <div class="bigbird">
          Hi, I'm Big Bird
        </div>
      </div>
    </div>
  </body>
</html>
`

const testSiblingHTML = `
<html>
   <body>
   		<div>
   			<p id="a" class="t1">
   				aaa
   				<br/>
   				<a class="t3">test anchor</a>
   			</p>
   			<p id="b" class="t2">bbb
   				<a class="t1">another anchor</a>
   			</p>
   			<p id="c" class="t3">ccc</p>
   			<p id="d" class="t4">ddd</p>
   		</div>
   	</body>
 </html>
 `

func TestFindAllNestedReturnsNestedMatchingNodes(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testHTML))
	allResults := htmls.FindAllNested(node, htmls.ByClass("bigbird"))

	if len(allResults) != 2 {
		t.Error("Expected 2 nodes returned but only found", len(allResults))
	}
}

func TestFindAllDoesNotReturnNestedMatchingNodes(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testHTML))
	allResults := htmls.FindAll(node, htmls.ByClass("bigbird"))

	if len(allResults) != 1 {
		t.Error("Expected 1 node returned but found", len(allResults))
	}
}

func TestFindNextSiblingReturnsMatchingNode(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testSiblingHTML))

	startingPoint, ok := htmls.Find(node, htmls.ByID("a"))
	if !ok {
		t.Error("Expected to find 'a' P node")
	} else {
		t3Node, ok := htmls.FindNextSibling(startingPoint, htmls.ByClass("t3"))
		if !ok || t3Node == nil {
			t.Error("Expected to find a node")
		} else {
			if t3Node.DataAtom != atom.P || htmls.Text(t3Node) != "ccc" {
				t.Error("Expected to find the third P node")

			}
		}
	}
}

func TestFindNextSiblingDoesntReturnChildren(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testSiblingHTML))

	startingPoint, ok := htmls.Find(node, htmls.ByID("b"))
	if !ok {
		t.Error("Expected to find 'b' P node")
	} else {
		_, ok := htmls.FindNextSibling(startingPoint, htmls.ByClass("t1"))
		if ok {
			t.Error("Didn't expect to find a next sibling node")
		}
	}
}

func TestFindPrevSiblingReturnsMatchingNode(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testSiblingHTML))

	startingPoint, ok := htmls.Find(node, htmls.ByID("c"))
	if !ok {
		t.Error("Expected to find the 'c' P node")
	} else {
		t1Node, ok := htmls.FindPrevSibling(startingPoint, htmls.ByClass("t1"))
		if !ok || t1Node == nil {
			t.Error("Expected to find a node")
		} else {
			if t1Node.DataAtom != atom.P || htmls.Text(t1Node) != "aaa test anchor" {
				t.Error("Expected to find the first P node")

			}
		}
	}
}

func TestFindPrevSiblingDoesntReturnChildren(t *testing.T) {
	node, _ := html.Parse(strings.NewReader(testSiblingHTML))

	startingPoint, ok := htmls.Find(node, htmls.ByID("c"))
	if !ok {
		t.Error("Expected to find 'c' P node")
	} else {
		_, ok := htmls.FindPrevSibling(startingPoint, htmls.ByClass("t3"))
		if ok {
			t.Error("Didn't expect to find a next sibling node")
		}
	}
}
