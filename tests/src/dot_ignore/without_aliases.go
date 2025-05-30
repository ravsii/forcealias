package alias

import (
	"fmt"
	. "io"
)

func withoutAliasFunc() {
	_ = fmt.Append

	fmt.Fprint(Discard, "nothing")
}
