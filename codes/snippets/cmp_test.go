// cmp example.
package snippets

// type Ordered interface { ... }

import (
	"cmp"
	"fmt"
	"testing"
)

func TestCmp(t *testing.T) {
	fmt.Println(cmp.Compare(1, 2))
	fmt.Println(cmp.Compare(2, 1))
	fmt.Println(cmp.Compare(1, 1))
}
