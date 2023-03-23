package builder

import (
	"fmt"
	"strings"
)

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.StringImpl(0)
}

func (e *HtmlElement) StringImpl(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indent*2)
	sb.WriteString(fmt.Sprintf("%s<%s>", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(e.text)
	}
	sb.WriteString(fmt.Sprintf("</%s>", e.name))
	sb.WriteString(" ")
	return sb.String()
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{rootName: rootName, root: HtmlElement{name: rootName}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
	return b
}

func Builder() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, w := range words {
		sb.WriteString("<li>")
		sb.WriteString(w)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// fluent builder
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())
}
