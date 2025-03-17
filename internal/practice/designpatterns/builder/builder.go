package builder

import (
	"fmt"
	"strings"
)

type HTMLBuilder struct {
	rootName string
	root     HTMLElement
}

type HTMLElement struct {
	name, text string
	elements   []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.StringImpl(0)
}

func (e *HTMLElement) StringImpl(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indent*2) //nolint:gomnd // indent is 2 spaces
	sb.WriteString(fmt.Sprintf("%s<%s>", i, e.name))

	if e.text != "" {
		sb.WriteString(e.text)
	}

	sb.WriteString(fmt.Sprintf("</%s>", e.name))
	sb.WriteString(" ")

	return sb.String()
}

func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{rootName: rootName, root: HTMLElement{name: rootName}}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

func (b *HTMLBuilder) AddChild(childName, childText string) {
	e := HTMLElement{name: childName, text: childText}
	b.root.elements = append(b.root.elements, e)
}

func (b *HTMLBuilder) AddChildFluent(childName, childText string) *HTMLBuilder {
	e := HTMLElement{name: childName, text: childText}
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

	sb.Reset()
	sb.WriteString("<ul>")

	words := []string{"hello", "world"}
	for _, w := range words {
		sb.WriteString("<li>")
		sb.WriteString(w)
		sb.WriteString("</li>")
	}

	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	// fluent builder
	b := NewHTMLBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())
}
