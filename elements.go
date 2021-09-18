package htmlgo

//     "a": HTMLAnchorElement;
func A(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("a").Children(children...)
}

//     "abbr": HTMLElement;
func Abbr(text string) (r *HTMLTagBuilder) {
	return Tag("abbr").Text(text)
}

//     "address": HTMLElement;
func Address(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("address").Children(children...)
}

//     "applet": HTMLAppletElement;
// Not support

//     "area": HTMLAreaElement;
func Area() (r *HTMLTagBuilder) {
	return Tag("area").OmitEndTag()
}

//     "article": HTMLElement;
func Article(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("article").Children(children...)
}

//     "aside": HTMLElement;
func Aside(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("aside").Children(children...)
}

//     "audio": HTMLAudioElement;
func Audio(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("audio").Children(children...)
}

//     "b": HTMLElement;
func B(text string) (r *HTMLTagBuilder) {
	return Tag("b").Text(text)
}

//     "base": HTMLBaseElement;
func Base() (r *HTMLTagBuilder) {
	return Tag("base").OmitEndTag()
}

//     "basefont": HTMLBaseFontElement;
// Not Support

//     "bdi": HTMLElement;
func Bdi(text string) (r *HTMLTagBuilder) {
	return Tag("bdi").Text(text)
}

//     "bdo": HTMLElement;
func Bdo(text string) (r *HTMLTagBuilder) {
	return Tag("bdo").Text(text)
}

//     "blockquote": HTMLQuoteElement;
func Blockquote(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("blockquote").Children(children...)
}

//     "body": HTMLBodyElement;
func Body(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("body").Children(children...)
}

//     "br": HTMLBRElement;
func Br() (r *HTMLTagBuilder) {
	return Tag("br").OmitEndTag()
}

//     "button": HTMLButtonElement;
func Button(label string) (r *HTMLTagBuilder) {
	return Tag("button").Text(label)
}

//     "canvas": HTMLCanvasElement;
func Canvas(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("canvas").Children(children...)
}

//     "caption": HTMLTableCaptionElement;
func Caption(text string) (r *HTMLTagBuilder) {
	return Tag("caption").Text(text)
}

//     "cite": HTMLElement;
func Cite(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("cite").Children(children...)
}

//     "code": HTMLElement;
func Code(text string) (r *HTMLTagBuilder) {
	return Tag("code").Text(text)
}

//     "col": HTMLTableColElement;
func Col() (r *HTMLTagBuilder) {
	return Tag("col").OmitEndTag()
}

//     "colgroup": HTMLTableColElement;
func Colgroup(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("colgroup").Children(children...)
}

//     "data": HTMLDataElement;
func Data(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("data").Children(children...)
}

//     "datalist": HTMLDataListElement;
func Datalist(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("datalist").Children(children...)
}

//     "dd": HTMLElement;
func Dd(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("dd").Children(children...)
}

//     "del": HTMLModElement;
func Del(text string) (r *HTMLTagBuilder) {
	return Tag("del").Text(text)
}

//     "details": HTMLDetailsElement;

func Details(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("details").Children(children...)
}

//     "dfn": HTMLElement;
func Dfn(text string) (r *HTMLTagBuilder) {
	return Tag("dfn").Text(text)
}

//     "dialog": HTMLDialogElement;
func Dialog(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("dialog").Children(children...)
}

//     "dir": HTMLDirectoryElement;
// Not Support

//     "div": HTMLDivElement;
func Div(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("div").Children(children...)
}

//     "dl": HTMLDListElement;
func Dl(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("dl").Children(children...)
}

//     "dt": HTMLElement;
func Dt(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("dt").Children(children...)
}

//     "em": HTMLElement;
func Em(text string) (r *HTMLTagBuilder) {
	return Tag("em").Text(text)
}

//     "embed": HTMLEmbedElement;
func Embed() (r *HTMLTagBuilder) {
	return Tag("embed").OmitEndTag()
}

//     "fieldset": HTMLFieldSetElement;
func Fieldset(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("fieldset").Children(children...)
}

//     "figcaption": HTMLElement;
func Figcaption(text string) (r *HTMLTagBuilder) {
	return Tag("figcaption").Text(text)
}

//     "figure": HTMLElement;
func Figure(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("figure").Children(children...)
}

//     "font": HTMLFontElement;
// Not Support

//     "footer": HTMLElement;
func Footer(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("footer").Children(children...)
}

//     "form": HTMLFormElement;
func Form(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("form").Children(children...)
}

//     "frame": HTMLFrameElement;
// Not Support

//     "frameset": HTMLFrameSetElement;
// Not Support

//     "h1": HTMLHeadingElement;
func H1(text string) (r *HTMLTagBuilder) {
	return Tag("h1").Text(text)
}

//     "h2": HTMLHeadingElement;
func H2(text string) (r *HTMLTagBuilder) {
	return Tag("h2").Text(text)
}

//     "h3": HTMLHeadingElement;
func H3(text string) (r *HTMLTagBuilder) {
	return Tag("h3").Text(text)
}

//     "h4": HTMLHeadingElement;
func H4(text string) (r *HTMLTagBuilder) {
	return Tag("h4").Text(text)
}

//     "h5": HTMLHeadingElement;
func H5(text string) (r *HTMLTagBuilder) {
	return Tag("h5").Text(text)
}

//     "h6": HTMLHeadingElement;
func H6(text string) (r *HTMLTagBuilder) {
	return Tag("h6").Text(text)
}

//     "head": HTMLHeadElement;

func Head(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("head").Children(children...)
}

//     "header": HTMLElement;
func Header(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("header").Children(children...)
}

//     "hgroup": HTMLElement;
func Hgroup(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("hgroup").Children(children...)
}

//     "hr": HTMLHRElement;
func Hr() (r *HTMLTagBuilder) {
	return Tag("hr").OmitEndTag()
}

//     "html": HTMLHtmlElement;
func HTML(children ...HTMLComponent) (r HTMLComponent) {
	return HTMLComponents{
		RawHTML("<!DOCTYPE html>\n"),
		Tag("html").Children(children...),
	}
}

//     "i": HTMLElement;
func I(text string) (r *HTMLTagBuilder) {
	return Tag("i").Text(text)
}

//     "iframe": HTMLIFrameElement;
func Iframe(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("iframe").Children(children...)
}

//     "img": HTMLImageElement;
func Img(src string) (r *HTMLTagBuilder) {
	return Tag("img").OmitEndTag().Attr("src", src)
}

//     "input": HTMLInputElement;

func Input(name string) (r *HTMLTagBuilder) {
	return Tag("input").OmitEndTag().Attr("name", name)
}

//     "ins": HTMLModElement;
func Ins(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ins").Children(children...)
}

//     "kbd": HTMLElement;
func Kbd(text string) (r *HTMLTagBuilder) {
	return Tag("kbd").Text(text)
}

//     "label": HTMLLabelElement;
func Label(text string) (r *HTMLTagBuilder) {
	return Tag("label").Text(text)
}

//     "legend": HTMLLegendElement;
func Legend(text string) (r *HTMLTagBuilder) {
	return Tag("legend").Text(text)
}

//     "li": HTMLLIElement;
func Li(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("li").Children(children...)
}

//     "link": HTMLLinkElement;
func Link(href string) (r *HTMLTagBuilder) {
	return Tag("link").OmitEndTag().Attr("href", href)
}

//     "main": HTMLElement;
func Main(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("main").Children(children...)
}

//     "map": HTMLMapElement;
func Map(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("map").Children(children...)
}

//     "mark": HTMLElement;
func Mark(text string) (r *HTMLTagBuilder) {
	return Tag("mark").Text(text)
}

//     "marquee": HTMLMarqueeElement;
// Not Support

//     "menu": HTMLMenuElement;
func Menu(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("menu").Children(children...)
}

//     "meta": HTMLMetaElement;

func Meta() (r *HTMLTagBuilder) {
	return Tag("meta").OmitEndTag()
}

//     "meter": HTMLMeterElement;
func Meter(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("meter").Children(children...)
}

//     "nav": HTMLElement;
func Nav(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("nav").Children(children...)
}

//     "noscript": HTMLElement;
func Noscript(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("noscript").Children(children...)
}

//     "object": HTMLObjectElement;
func Object(data string) (r *HTMLTagBuilder) {
	return Tag("object").Attr("data", data)
}

//     "ol": HTMLOListElement;
func Ol(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ol").Children(children...)
}

//     "optgroup": HTMLOptGroupElement;
func Optgroup(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("optgroup").Children(children...)
}

//     "option": HTMLOptionElement;
func Option(text string) (r *HTMLTagBuilder) {
	return Tag("option").Text(text)
}

//     "output": HTMLOutputElement;
func Output(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("output").Children(children...)
}

//     "p": HTMLParagraphElement;
func P(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("p").Children(children...)
}

//     "param": HTMLParamElement;
func Param(name string) (r *HTMLTagBuilder) {
	return Tag("param").OmitEndTag().Attr("name", name)
}

//     "picture": HTMLPictureElement;
func Picture(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("picture").Children(children...)
}

//     "pre": HTMLPreElement;
func Pre(text string) (r *HTMLTagBuilder) {
	return Tag("pre").Text(text)
}

//     "progress": HTMLProgressElement;
func Progress(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("progress").Children(children...)
}

//     "q": HTMLQuoteElement;
func Q(text string) (r *HTMLTagBuilder) {
	return Tag("q").Text(text)
}

//     "rp": HTMLElement;
func Rp(text string) (r *HTMLTagBuilder) {
	return Tag("rp").Text(text)
}

//     "rt": HTMLElement;
func Rt(text string) (r *HTMLTagBuilder) {
	return Tag("rt").Text(text)
}

//     "ruby": HTMLElement;
func Ruby(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ruby").Children(children...)
}

//     "s": HTMLElement;
func S(text string) (r *HTMLTagBuilder) {
	return Tag("s").Text(text)
}

//     "samp": HTMLElement;

func Samp(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("samp").Children(children...)
}

//     "script": HTMLScriptElement;
func Script(script string) (r *HTMLTagBuilder) {
	return Tag("script").
		Attr("type", "text/javascript").
		Children(RawHTML(script))
}

//     "section": HTMLElement;
func Section(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("section").Children(children...)
}

//     "select": HTMLSelectElement;
func Select(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("select").Children(children...)
}

//     "slot": HTMLSlotElement;
func Slot(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("slot").Children(children...)
}

//     "small": HTMLElement;
func Small(text string) (r *HTMLTagBuilder) {
	return Tag("small").Text(text)
}

//     "source": HTMLSourceElement;
func Source(src string) (r *HTMLTagBuilder) {
	return Tag("source").OmitEndTag().Attr("src", src)
}

//     "span": HTMLSpanElement;
func Span(text string) (r *HTMLTagBuilder) {
	return Tag("span").Text(text)
}

//     "strong": HTMLElement;
func Strong(text string) (r *HTMLTagBuilder) {
	return Tag("strong").Text(text)
}

//     "style": HTMLStyleElement;
func Style(style string) (r *HTMLTagBuilder) {
	return Tag("style").
		Attr("type", "text/css").
		Children(RawHTML(style))
}

//     "sub": HTMLElement;
func Sub(text string) (r *HTMLTagBuilder) {
	return Tag("sub").Text(text)
}

//     "summary": HTMLElement;
func Summary(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("summary").Children(children...)
}

//     "sup": HTMLElement;
func Sup(text string) (r *HTMLTagBuilder) {
	return Tag("sup").Text(text)
}

//     "table": HTMLTableElement;
func Table(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("table").Children(children...)
}

//     "tbody": HTMLTableSectionElement;
func Tbody(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("tbody").Children(children...)
}

//     "td": HTMLTableDataCellElement;
func Td(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("td").Children(children...)
}

//     "template": HTMLTemplateElement;
func Template(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("template").Children(children...)
}

//     "textarea": HTMLTextAreaElement;
func Textarea(text string) (r *HTMLTagBuilder) {
	return Tag("textarea").Text(text)
}

//     "tfoot": HTMLTableSectionElement;
func Tfoot(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("tfoot").Children(children...)
}

//     "th": HTMLTableHeaderCellElement;
func Th(text string) (r *HTMLTagBuilder) {
	return Tag("th").Text(text)
}

//     "thead": HTMLTableSectionElement;
func Thead(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("thead").Children(children...)
}

//     "time": HTMLTimeElement;
func Time(datetime string) (r *HTMLTagBuilder) {
	return Tag("time").Attr("datetime", datetime)
}

//     "title": HTMLTitleElement;
func Title(text string) (r *HTMLTagBuilder) {
	return Tag("title").Text(text)
}

//     "tr": HTMLTableRowElement;
func Tr(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("tr").Children(children...)
}

//     "track": HTMLTrackElement;
func Track(src string) (r *HTMLTagBuilder) {
	return Tag("track").OmitEndTag().Attr("src", src)
}

//     "u": HTMLElement;
func U(text string) (r *HTMLTagBuilder) {
	return Tag("u").Text(text)
}

//     "ul": HTMLUListElement;
func Ul(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("ul").Children(children...)
}

//     "var": HTMLElement;
func Var(text string) (r *HTMLTagBuilder) {
	return Tag("var").Text(text)
}

//     "video": HTMLVideoElement;
func Video(children ...HTMLComponent) (r *HTMLTagBuilder) {
	return Tag("video").Children(children...)
}

//     "wbr": HTMLElement;
func Wbr() (r *HTMLTagBuilder) {
	return Tag("wbr").OmitEndTag()
}
