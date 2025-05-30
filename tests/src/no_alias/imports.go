package alias

import (
	"fmt"
	_ "fmt"
	x "fmt"
	. "io"
)

func goplsDontRemove() {
	_ = fmt.Append
	_ = x.Append

	x.Fprint(Discard, "nothing")
}
