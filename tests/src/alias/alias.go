package alias

import (
	"fmt"
	x "fmt"
	"net/url"
)

func goplsDontRemove() {
	_ = fmt.Append
	_ = x.Append
	_ = url.URL{}
}
