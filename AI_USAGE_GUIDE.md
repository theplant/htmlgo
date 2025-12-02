# htmlgo Library - AI Usage Guide

## Overview
**htmlgo** is a type-safe Go library for generating HTML on the server side. It provides a fluent, builder-pattern API for constructing HTML components programmatically.

**Package**: `github.com/theplant/htmlgo`

## Core Concepts

### 1. HTMLComponent Interface
All HTML elements implement this interface:
```go
type HTMLComponent interface {
    MarshalHTML(ctx context.Context) ([]byte, error)
}
```

### 2. Import Pattern
Recommended to use dot import for cleaner code:
```go
import . "github.com/theplant/htmlgo"
```

## Key Functions & Types

### Basic Building Blocks

**Text Rendering** (auto-escaped):
- `Text(string)` - Escapes HTML entities
- `Textf(format, ...args)` - Formatted text with escaping
- `RawHTML(string)` - Unescaped HTML string

**Output**:
- `Fprint(w io.Writer, root HTMLComponent, ctx context.Context)` - Write to writer
- `MustString(root HTMLComponent, ctx context.Context)` - Convert to string (panics on error)

### HTML Elements

**Container Elements** (accept children):
```go
Div(...children)
Span(text)
A(...children)
Body(...children)
Head(...children)
Section(...children)
Article(...children)
Header(...children)
Footer(...children)
Nav(...children)
Ul(...children), Ol(...children), Li(...children)
Table(...children), Thead(...children), Tbody(...children), Tr(...children), Td(...children), Th(text)
Form(...children)
Fieldset(...children)
Select(...children), Option(text)
```

**Text Elements** (accept text):
```go
H1(text), H2(text), H3(text), H4(text), H5(text), H6(text)
P(...children)
Button(label)
Label(text)
Textarea(text)
Strong(text), Em(text), B(text), I(text)
Code(text), Pre(text)
```

**Self-Closing Elements**:
```go
Br()
Hr()
Img(src)
Input(name)
Meta()
Link(href)
```

**Special Elements**:
- `HTML(...children)` - Adds `<!DOCTYPE html>` automatically
- `Script(jsCode)` - Wraps in script tag with type="text/javascript"
- `Style(cssCode)` - Wraps in style tag with type="text/css"

### HTMLTagBuilder Methods

**Attributes**:
```go
.Attr(key, value, ...)           // Set any attribute (variadic pairs)
.AttrIf(key, value, condition)   // Conditional attribute
.Id(string)
.Class(...string)                // Add classes (space-separated supported)
.ClassIf(string, bool)           // Conditional class
.Data(key, value, ...)           // data-* attributes
.Href(string), .Src(string), .Alt(string)
.Name(string), .Value(string), .Type(string)
.Placeholder(string), .Title(string)
.Action(string), .Method(string)
.For(string), .Role(string), .Target(string)
.Disabled(bool), .Checked(bool), .Required(bool), .Readonly(bool)
.TabIndex(int)
```

**Styling**:
```go
.Style(string)              // Add inline styles (semicolon-separated)
.StyleIf(string, bool)      // Conditional styles
```

**Content Manipulation**:
```go
.Text(string)                    // Set text content (escaped)
.Children(...HTMLComponent)      // Replace children
.AppendChildren(...HTMLComponent)
.PrependChildren(...HTMLComponent)
```

**Advanced**:
```go
.SetAttr(key, value)        // Programmatically set attribute
.Tag(string)                // Set tag name
.OmitEndTag()              // For self-closing tags
```

### Attribute Value Types
Attributes accept multiple types:
- `string`, `[]byte`, `[]rune` - Used as-is
- `int`, `uint`, `float` types - Converted to string
- `bool` - If true, renders as boolean attribute; if false, omitted
- Other types - JSON-encoded

### Custom Components

**ComponentFunc**:
```go
ComponentFunc(func(ctx context.Context) ([]byte, error))
```

**Custom Builder Pattern**:
```go
type MyBuilder struct {
    field1 string
    field2 int
}

func (b *MyBuilder) Field1(v string) *MyBuilder {
    b.field1 = v
    return b
}

func (b *MyBuilder) MarshalHTML(ctx context.Context) ([]byte, error) {
    return Div(Text(b.field1)).MarshalHTML(ctx)
}
```

### Conditional Rendering

**If/ElseIf/Else** (with pre-evaluated components):
```go
If(condition, component1, component2).
    ElseIf(condition2, component3).
    Else(component4)
```

**Iff/ElseIf/Else** (with lazy evaluation):
```go
Iff(condition, func() HTMLComponent {
    return Div(Text("Lazy evaluated"))
}).ElseIf(condition2, func() HTMLComponent {
    return Span("Alternative")
}).Else(func() HTMLComponent {
    return Text("Default")
})
```

**Use Iff when**: Body depends on condition or expensive to compute.

### Context Usage

Pass data through context:
```go
ctx := context.WithValue(context.TODO(), "user", userData)

ComponentFunc(func(ctx context.Context) ([]byte, error) {
    if user, ok := ctx.Value("user").(*User); ok {
        return Div(Text(user.Name)).MarshalHTML(ctx)
    }
    return nil, nil
})
```

### Component Composition

**HTMLComponents** (multiple components):
```go
Components(comp1, comp2, comp3)
```

**Layout Pattern**:
```go
func layout(content HTMLComponent) HTMLComponent {
    return HTML(
        Head(Meta().Charset("utf8")),
        Body(header(), content, footer()),
    )
}
```

## Common Patterns

### 1. Simple Page
```go
HTML(
    Head(Title("Page")),
    Body(Div(Text("Content")))
)
```

### 2. Form with Validation
```go
Form(
    Input("email").Type("email").Required(true),
    Button("Submit").Type("submit"),
).Action("/submit").Method("post")
```

### 3. Dynamic List
```go
items := []string{"A", "B", "C"}
children := make([]HTMLComponent, len(items))
for i, item := range items {
    children[i] = Li(Text(item))
}
Ul(children...)
```

### 4. Conditional Classes
```go
Div().
    Class("base-class").
    ClassIf("active", isActive).
    ClassIf("error", hasError)
```

### 5. Data Attributes
```go
Div().Data("user-id", "123", "role", "admin")
// Renders: <div data-user-id='123' data-role='admin'>
```

### 6. HTTP Handler Integration
```go
func handler(w http.ResponseWriter, r *http.Request) {
    user := getUserFromSession(r)
    ctx := context.WithValue(context.TODO(), "user", user)
    
    page := HTML(
        Head(Title("My Page")),
        Body(Div(Text("Hello"))),
    )
    
    Fprint(w, page, ctx)
}
```

## Important Notes

- **Text is auto-escaped** via `Text()` - use `RawHTML()` for unescaped content
- **Context is required** for MarshalHTML - use `context.TODO()` if not needed
- **Nil children are filtered** automatically
- **Boolean attributes**: `true` renders attribute, `false` omits it
- **Styles are concatenated** with semicolons automatically
- **Classes are space-split** and deduplicated
- **Attributes use single quotes** in output: `<div class='myclass'>`

## Performance
- Uses `sync.Pool` for buffer reuse
- Efficient string building with `bytes.Buffer`
- Minimal allocations for common operations

## Quick Reference

| Task | Code |
|------|------|
| Escaped text | `Text("content")` |
| Raw HTML | `RawHTML("<svg>...</svg>")` |
| Container | `Div(child1, child2)` |
| Attributes | `.Attr("key", "value")` |
| Classes | `.Class("class1 class2")` |
| Styles | `.Style("color:red; font-size:14px")` |
| Conditional | `Iff(cond, func() HTMLComponent {...})` |
| Custom component | `ComponentFunc(func(ctx) ([]byte, error) {...})` |
| Output | `Fprint(w, component, ctx)` |