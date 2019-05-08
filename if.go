package htmlgo

import "context"

type IfBuilder struct {
	comps []HTMLComponent
	set   bool
}

func If(v bool, comps ...HTMLComponent) (r *IfBuilder) {
	r = &IfBuilder{}
	if v {
		r.comps = comps
		r.set = true
	}
	return
}

func (b *IfBuilder) ElseIf(v bool, comps ...HTMLComponent) (r *IfBuilder) {
	if b.set {
		return b
	}
	if v {
		b.comps = comps
		b.set = true
	}
	return b
}

func (b *IfBuilder) Else(comps ...HTMLComponent) (r *IfBuilder) {
	if b.set {
		return b
	}
	b.set = true
	b.comps = comps
	return b
}

func (b *IfBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if len(b.comps) == 0 {
		return
	}
	return HTMLComponents(b.comps).MarshalHTML(ctx)
}
