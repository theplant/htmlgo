/*

## htmlgo

Type safe and modularize way to generate html on server side.

Import the package with `.` gives you simpler code:

	import (
		. "github.com/theplant/htmlgo"
	)

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
