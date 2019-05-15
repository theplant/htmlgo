package htmlgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"strings"
)

type tagAttr struct {
	key   string
	value interface{}
}

type HTMLTagBuilder struct {
	tag        string
	attrs      []*tagAttr
	classNames []string
	children   []HTMLComponent
}

func Tag(tag string) (r *HTMLTagBuilder) {
	r = &HTMLTagBuilder{}

	if r.attrs == nil {
		r.attrs = []*tagAttr{}
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

func (b *HTMLTagBuilder) SetAttr(k string, v interface{}) {
	for _, at := range b.attrs {
		if at.key == k {
			at.value = v
			return
		}
	}
	b.attrs = append(b.attrs, &tagAttr{k, v})
}

func (b *HTMLTagBuilder) Attr(vs ...interface{}) (r *HTMLTagBuilder) {
	if len(vs)%2 != 0 {
		vs = append(vs, "")
	}

	for i := 0; i < len(vs); i = i + 2 {
		if key, ok := vs[i].(string); ok {
			b.SetAttr(key, vs[i+1])
		} else {
			panic(fmt.Sprintf("Attr key must be string, but was %#+v", vs[i]))
		}
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
	b.Attr("tabindex", v)
	return b
}

func (b *HTMLTagBuilder) Required(v bool) (r *HTMLTagBuilder) {
	b.Attr("required", v)
	return b
}

func (b *HTMLTagBuilder) Readonly(v bool) (r *HTMLTagBuilder) {
	b.Attr("readonly", v)
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
	b.Attr("disabled", v)
	return b
}

func (b *HTMLTagBuilder) Checked(v bool) (r *HTMLTagBuilder) {
	b.Attr("checked", v)
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
		var seg string
		switch v := at.value.(type) {
		case string:
			seg = fmt.Sprintf("%s='%s'", at.key, v)
		case []byte:
			seg = fmt.Sprintf("%s='%s'", at.key, string(v))
		case []rune:
			seg = fmt.Sprintf("%s='%s'", at.key, string(v))
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			seg = fmt.Sprintf("%s='%d'", at.key, v)
		case float32, float64:
			seg = fmt.Sprintf("%s='%f'", at.key, v)
		case bool:
			if v {
				seg = at.key
			}
		default:
			bseg, _ := json.Marshal(v)
			val := strings.ReplaceAll(string(bseg), "\n", "")
			seg = fmt.Sprintf("%s='%s'", at.key, val)
		}
		attrSegs = append(attrSegs, seg)
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
