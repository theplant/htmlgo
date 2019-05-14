package htmlgo

import (
	"bytes"
	"context"
	"fmt"
	"html"
	"strings"
)

type HTMLTagBuilder struct {
	tag        string
	attrs      [][]string
	classNames []string
	children   []HTMLComponent
}

func Tag(tag string) (r *HTMLTagBuilder) {
	r = &HTMLTagBuilder{}

	if r.attrs == nil {
		r.attrs = [][]string{}
	}

	r.Tag(tag)

	return
}

func (b *HTMLTagBuilder) Tag(v string) (r *HTMLTagBuilder) {
	b.tag = v
	return b
}

func (b *HTMLTagBuilder) Text(v string) (r *HTMLTagBuilder) {
	b.Children(Text(v))
	return b
}

func (b *HTMLTagBuilder) Children(comps ...HTMLComponent) (r *HTMLTagBuilder) {
	b.children = comps
	return b
}

func (b *HTMLTagBuilder) SetAttr(k string, v string) {
	for _, at := range b.attrs {
		if at[0] == k {
			at[1] = v
			return
		}
	}
	b.attrs = append(b.attrs, []string{k, v})
}

func (b *HTMLTagBuilder) Attr(vs ...string) (r *HTMLTagBuilder) {
	if len(vs)%2 != 0 {
		vs = append(vs, "")
	}

	for i := 0; i < len(vs); i = i + 2 {
		b.SetAttr(vs[i], vs[i+1])
	}
	return b
}

func (b *HTMLTagBuilder) Class(names ...string) (r *HTMLTagBuilder) {
	b.addClass(names...)
	return b
}

func (b *HTMLTagBuilder) addClass(names ...string) (r *HTMLTagBuilder) {
	for _, n := range names {
		ins := strings.Split(n, " ")
		for _, in := range ins {
			tin := strings.TrimSpace(in)
			if len(tin) > 0 {
				b.classNames = append(b.classNames, tin)
			}
		}
	}
	return b
}

func (b *HTMLTagBuilder) ClassIf(name string, add bool) (r *HTMLTagBuilder) {
	if !add {
		return b
	}
	b.addClass(name)
	return b
}

func (b *HTMLTagBuilder) Data(vs ...string) (r *HTMLTagBuilder) {
	for i := 0; i < len(vs); i = i + 2 {
		b.Attr(fmt.Sprintf("data-%s", vs[i]), vs[i+1])
	}
	return b
}

func (b *HTMLTagBuilder) Id(v string) (r *HTMLTagBuilder) {
	b.Attr("id", v)
	return b
}

func (b *HTMLTagBuilder) Href(v string) (r *HTMLTagBuilder) {
	b.Attr("href", v)
	return b
}

func (b *HTMLTagBuilder) Rel(v string) (r *HTMLTagBuilder) {
	b.Attr("ref", v)
	return b
}

func (b *HTMLTagBuilder) Title(v string) (r *HTMLTagBuilder) {
	b.Attr("title", html.EscapeString(v))
	return b
}

func (b *HTMLTagBuilder) TabIndex(v int) (r *HTMLTagBuilder) {
	b.Attr("tabindex", fmt.Sprint(v))
	return b
}

func (b *HTMLTagBuilder) Required(v bool) (r *HTMLTagBuilder) {
	b.Attr("required", fmt.Sprint(v))
	return b
}

func (b *HTMLTagBuilder) Readonly(v bool) (r *HTMLTagBuilder) {
	b.Attr("readonly", fmt.Sprint(v))
	return b
}

func (b *HTMLTagBuilder) Role(v string) (r *HTMLTagBuilder) {
	b.Attr("role", v)
	return b
}

func (b *HTMLTagBuilder) Alt(v string) (r *HTMLTagBuilder) {
	b.Attr("alt", v)
	return b
}

func (b *HTMLTagBuilder) Target(v string) (r *HTMLTagBuilder) {
	b.Attr("target", v)
	return b
}

func (b *HTMLTagBuilder) Name(v string) (r *HTMLTagBuilder) {
	b.Attr("name", v)
	return b
}

func (b *HTMLTagBuilder) Value(v string) (r *HTMLTagBuilder) {
	b.Attr("value", v)
	return b
}

func (b *HTMLTagBuilder) Style(v string) (r *HTMLTagBuilder) {
	b.Attr("style", v)
	return b
}

func (b *HTMLTagBuilder) Type(v string) (r *HTMLTagBuilder) {
	b.Attr("type", v)
	return b
}

func (b *HTMLTagBuilder) Placeholder(v string) (r *HTMLTagBuilder) {
	b.Attr("placeholder", v)
	return b
}

func (b *HTMLTagBuilder) Src(v string) (r *HTMLTagBuilder) {
	b.Attr("src", v)
	return b
}

func (b *HTMLTagBuilder) Property(v string) (r *HTMLTagBuilder) {
	b.Attr("property", v)
	return b
}

func (b *HTMLTagBuilder) Action(v string) (r *HTMLTagBuilder) {
	b.Attr("action", v)
	return b
}

func (b *HTMLTagBuilder) Method(v string) (r *HTMLTagBuilder) {
	b.Attr("method", v)
	return b
}

func (b *HTMLTagBuilder) Content(v string) (r *HTMLTagBuilder) {
	b.Attr("content", v)
	return b
}

func (b *HTMLTagBuilder) Charset(v string) (r *HTMLTagBuilder) {
	b.Attr("charset", v)
	return b
}

func (b *HTMLTagBuilder) Disabled(v bool) (r *HTMLTagBuilder) {
	b.Attr("disabled", fmt.Sprint(v))
	return b
}

func (b *HTMLTagBuilder) AppendChildren(c ...HTMLComponent) (r *HTMLTagBuilder) {
	b.children = append(b.children, c...)
	return b
}

func (b *HTMLTagBuilder) PrependChildren(c ...HTMLComponent) (r *HTMLTagBuilder) {
	b.children = append(c, b.children...)
	return b
}

func (b *HTMLTagBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	class := strings.TrimSpace(strings.Join(b.classNames, " "))
	if len(class) > 0 {
		b.Attr("class", class)
	}

	// remove empty
	cs := []HTMLComponent{}
	for _, c := range b.children {
		if c == nil {
			continue
		}
		cs = append(cs, c)
	}

	attrSegs := []string{}
	for _, at := range b.attrs {
		attrSegs = append(attrSegs, fmt.Sprintf("%s='%s'", at[0], at[1]))
	}

	attrStr := ""
	if len(attrSegs) > 0 {
		attrStr = " " + strings.Join(attrSegs, " ")
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("\n<%s%s>", b.tag, attrStr))
	if len(cs) > 0 {
		// buf.WriteString("\n")
		for _, c := range cs {
			var child []byte
			child, err = c.MarshalHTML(ctx)
			if err != nil {
				return
			}
			buf.Write(child)
		}
	}
	buf.WriteString(fmt.Sprintf("</%s>\n", b.tag))
	r = buf.Bytes()
	return
}
