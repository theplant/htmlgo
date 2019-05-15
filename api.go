/*

## htmlgo

Type safe and modularize way to generate html on server side.
Download the package with `go get -v github.com/theplant/htmlgo` and import the package with `.` gives you simpler code:

	import (
		. "github.com/theplant/htmlgo"
	)

also checkout full API documentation at: https://godoc.org/github.com/theplant/htmlgo

*/
package htmlgo

import (
	"context"
)

type HTMLComponent interface {
	MarshalHTML(ctx context.Context) ([]byte, error)
}

type ComponentFunc func(ctx context.Context) (r []byte, err error)

func (f ComponentFunc) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return f(ctx)
}

type MutableAttrHTMLComponent interface {
	HTMLComponent
	SetAttr(k string, v interface{})
}
