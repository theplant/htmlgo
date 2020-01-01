package htmlgo

import (
	"bytes"
	"context"
	"fmt"
	"html"
	"io"
)

type RawHTML string

func (s RawHTML) MarshalHTML(ctx context.Context) (r []byte, err error) {
	r = []byte(s)
	return
}

func Text(text string) (r HTMLComponent) {
	return RawHTML(html.EscapeString(text))
}

func Textf(format string, a ...interface{}) (r HTMLComponent) {
	return Text(fmt.Sprintf(format, a...))
}

type HTMLComponents []HTMLComponent

func Components(comps ...HTMLComponent) HTMLComponents {
	return HTMLComponents(comps)
}

func (hcs HTMLComponents) MarshalHTML(ctx context.Context) (r []byte, err error) {
	buf := bytes.NewBuffer(nil)
	for _, h := range hcs {
		if h == nil {
			continue
		}
		var b []byte
		b, err = h.MarshalHTML(ctx)
		if err != nil {
			return
		}
		buf.Write(b)
	}
	r = buf.Bytes()
	return
}

func Fprint(w io.Writer, root HTMLComponent, ctx context.Context) (err error) {
	if root == nil {
		return
	}
	var b []byte
	b, err = root.MarshalHTML(ctx)
	if err != nil {
		return
	}
	_, err = fmt.Fprint(w, string(b))
	return
}

func MustString(root HTMLComponent, ctx context.Context) string {
	b := bytes.NewBuffer(nil)
	err := Fprint(b, root, ctx)
	if err != nil {
		panic(err)
	}
	return b.String()
}
