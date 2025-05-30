package alias

import (
	_ "encoding/json" // want "should be aliased as notUnderscore"
	"fmt"             // want "should be aliased as testAlias"
	. "net/url"       // want "should be aliased as notDot"
)

func init() {
	_ = fmt.Append
	_ = URL{}
	// _ = Main // from singlechecked
}
