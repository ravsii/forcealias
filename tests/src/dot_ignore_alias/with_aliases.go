package alias

import (
	"fmt" // want "should be aliased as testAlias"
	. "io"
)

func withAliasFunc() {
	_ = fmt.Append

	fmt.Fprint(Discard, "nothing")
}
