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

type IfFuncBuilder struct {
	f   func() HTMLComponent
	set bool
}

func Iff(v bool, f func() HTMLComponent) (r *IfFuncBuilder) {
	r = &IfFuncBuilder{}
	if v {
		r.f = f
		r.set = true
	}
	return
}

func (b *IfFuncBuilder) ElseIf(v bool, f func() HTMLComponent) (r *IfFuncBuilder) {
	if b.set {
		return b
	}
	if v {
		b.f = f
		b.set = true
	}
	return b
}

func (b *IfFuncBuilder) Else(f func() HTMLComponent) (r *IfFuncBuilder) {
	if b.set {
		return b
	}
	b.set = true
	b.f = f
	return b
}

func (b *IfFuncBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	if b.f == nil {
		return
	}
	return b.f().MarshalHTML(ctx)
}
