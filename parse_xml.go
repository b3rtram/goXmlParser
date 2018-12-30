package goXmlParser

import (
	"bufio"
	"fmt"
	"io"
)

//Attribute
type Attribute struct {
	key   string
	value string
}

//Tag
type Tag struct {
	name       string
	attributes []Attribute
}

//Parse start parsing
func ParseChan(xml io.Reader, startTag chan Tag, endTag chan Tag) {

	Parse(xml,
		func(t Tag) {
			startTag <- t
		},
		func(t Tag) {
			endTag <- t
		})
}

//Parse start parsing
func Parse(xml io.Reader, startTag func(Tag), endTag func(Tag)) {

	r := bufio.NewReader(xml)
	for {

		b, err := r.ReadBytes('>')
		if err != nil {
			fmt.Println(err)
			break
		}

		i := 0
		for i < len(b) {

			if b[i] == '<' && b[i+1] != '/' {
				i++
				t := parseTag(&b, &i)
				startTag(t)
			} else {
				t := parseTag(&b, &i)
				endTag(t)
			}

			i++
		}
	}
}

func parseTag(b *[]byte, i *int) Tag {

	t := Tag{name: ""}
	for *i < len(*b) {
		if (*b)[*i] == ' ' || (*b)[*i] == '>' {
			*i++
			break
		}

		t.name += string((*b)[*i])
		*i++
	}

	a := Attribute{key: "", value: ""}

	for *i < len(*b) {

		buf := (*b)[*i]

		if buf == ' ' {
			*i++
			continue
		}

		if buf == '>' {
			break
		}

		if buf == '=' {
			*i += 2

			for *i < len(*b) {
				a.value += string((*b)[*i])
				if (*b)[*i] == '"' || (*b)[*i] == '\'' || (*b)[*i+1] == '"' || (*b)[*i+1] == '\'' {
					*i += 2
					break
				}

				*i++
			}

			t.attributes = append(t.attributes, a)
			a = Attribute{key: "", value: ""}

		} else {

			a.key += string(buf)
			*i++

		}

	}

	return t
}
