

## htmlgo

Type safe and modularize way to generate html on server side.
Download the package with `go get -v github.com/theplant/htmlgo` and import the package with `.` gives you simpler code:


```go
import (
	. "github.com/theplant/htmlgo"
)
```

also checkout full API documentation at: <a href="https://godoc.org/github.com/theplant/htmlgo">https://godoc.org/github.com/theplant/htmlgo</a>



Create a simple div, Text will be escaped by html
```go
	comp := Div(
	    Text("123<h1>"),
	)
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <div>123&lt;h1&gt;</div>
```

Create a full html page
```go
	comp := HTML(
	    Head(
	        Meta().Charset("utf8"),
	        Title("My test page"),
	    ),
	    Body(
	        Img("images/firefox-icon.png").Alt("My test image"),
	    ),
	)
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <!DOCTYPE html>
	//
	// <html>
	// <head>
	// <meta charset='utf8'></meta>
	//
	// <title>My test page</title>
	// </head>
	//
	// <body>
	// <img src='images/firefox-icon.png' alt='My test image'></img>
	// </body>
	// </html>
```

Use RawHTML and Component
```go
	userProfile := func(username string, avatarURL string) HTMLComponent {
	    return ComponentFunc(func(ctx context.Context) (r []byte, err error) {
	        return Div(
	            H1(username).Class("profileName"),
	            Img(avatarURL).Class("profileImage"),
	            RawHTML("<svg>complicated svg</svg>\n"),
	        ).Class("userProfile").MarshalHTML(ctx)
	    })
	}
	
	comp := Ul(
	    Li(
	        userProfile("felix<h1>", "http://image.com/img1.png"),
	    ),
	    Li(
	        userProfile("john", "http://image.com/img2.png"),
	    ),
	)
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <ul>
	// <li>
	// <div class='userProfile'>
	// <h1 class='profileName'>felix&lt;h1&gt;</h1>
	//
	// <img src='http://image.com/img1.png' class='profileImage'></img>
	// <svg>complicated svg</svg>
	// </div>
	// </li>
	//
	// <li>
	// <div class='userProfile'>
	// <h1 class='profileName'>john</h1>
	//
	// <img src='http://image.com/img2.png' class='profileImage'></img>
	// <svg>complicated svg</svg>
	// </div>
	// </li>
	// </ul>
```

More complicated customized component
```go
	/*
	    Define MySelect as follows:
	
	    type MySelectBuilder struct {
	        options  [][]string
	        selected string
	    }
	
	    func MySelect() *MySelectBuilder {
	        return &MySelectBuilder{}
	    }
	
	    func (b *MySelectBuilder) Options(opts [][]string) (r *MySelectBuilder) {
	        b.options = opts
	        return b
	    }
	
	    func (b *MySelectBuilder) Selected(selected string) (r *MySelectBuilder) {
	        b.selected = selected
	        return b
	    }
	
	    func (b *MySelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	        opts := []HTMLComponent{}
	        for _, op := range b.options {
	            var opt HTMLComponent
	            if op[0] == b.selected {
	                opt = Option(op[1]).Value(op[0]).Attr("selected", "true")
	            } else {
	                opt = Option(op[1]).Value(op[0])
	            }
	            opts = append(opts, opt)
	        }
	        return Select(opts...).MarshalHTML(ctx)
	    }
	*/
	
	comp := MySelect().Options([][]string{
	    {"1", "label 1"},
	    {"2", "label 2"},
	    {"3", "label 3"},
	}).Selected("2")
	
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <select>
	// <option value='1'>label 1</option>
	//
	// <option value='2' selected='true'>label 2</option>
	//
	// <option value='3'>label 3</option>
	// </select>
```

Write a little bit of JavaScript and stylesheet
```go
	comp := Div(
	    Button("Hello").Id("hello"),
	    Style(`
	.container {
	    background-color: red;
	}
	`),
	
	    Script(`
	var b = document.getElementById("hello")
	b.onclick = function(e){
	    alert("Hello");
	}
	`),
	).Class("container")
	
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <div class='container'>
	// <button id='hello'>Hello</button>
	//
	// <style type='text/css'>
	// 	.container {
	// 		background-color: red;
	// 	}
	// </style>
	//
	// <script type='text/javascript'>
	// 	var b = document.getElementById("hello")
	// 	b.onclick = function(e){
	// 		alert("Hello");
	// 	}
	// </script>
	// </div>
```

An example about how to integrate into http.Handler, and how to do layout, and how to use context.
```go
	type User struct {
	    Name string
	}
	
	userStatus := func() HTMLComponent {
	    return ComponentFunc(func(ctx context.Context) (r []byte, err error) {
	
	        if currentUser, ok := ctx.Value("currentUser").(*User); ok {
	            return Div(
	                Text(currentUser.Name),
	            ).Class("username").MarshalHTML(ctx)
	        }
	
	        return Div(Text("Login")).Class("login").MarshalHTML(ctx)
	    })
	}
	
	myHeader := func() HTMLComponent {
	    return Div(
	        Text("header"),
	        userStatus(),
	    ).Class("header")
	}
	myFooter := func() HTMLComponent {
	    return Div(Text("footer")).Class("footer")
	}
	
	layout := func(in HTMLComponent) (out HTMLComponent) {
	    out = HTML(
	        Head(
	            Meta().Charset("utf8"),
	        ),
	        Body(
	            myHeader(),
	            in,
	            myFooter(),
	        ),
	    )
	    return
	}
	
	getLoginUserFromCookie := func(r *http.Request) *User {
	    return &User{Name: "felix"}
	}
	
	homeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    user := getLoginUserFromCookie(r)
	    ctx := context.WithValue(context.TODO(), "currentUser", user)
	
	    root := Div(
	        Text("This is my home page"),
	    )
	
	    Fprint(w, layout(root), ctx)
	})
	
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	homeHandler.ServeHTTP(w, r)
	
	fmt.Println(w.Body.String())
	
	//Output:
	// <!DOCTYPE html>
	//
	// <html>
	// <head>
	// <meta charset='utf8'></meta>
	// </head>
	//
	// <body>
	// <div class='header'>header
	// <div class='username'>felix</div>
	// </div>
	//
	// <div>This is my home page</div>
	//
	// <div class='footer'>footer</div>
	// </body>
	// </html>
```

An example show how to set different type of attributes
```go
	type MoreData struct {
	    Name  string
	    Count int
	}
	comp := Div(
	    Input("username").
	        Type("checkbox").
	        Attr("checked", true).
	        Attr("more-data", &MoreData{Name: "felix", Count: 100}).
	        Attr("max-length", 10),
	)
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <div>
	// <input name='username' type='checkbox' checked more-data='{"Name":"felix","Count":100}' max-length='10'></input>
	// </div>
```



