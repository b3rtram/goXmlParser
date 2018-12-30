package goXmlParser

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {

	str := "<node id=\"6145450543\" lat=\"49.4569977\" lon=\"11.0380648\" version=\"1\" timestamp=\"2018-12-18T15:33:47Z\" changeset=\"0\"><tag k=\"addr:city\" v=\"Nünberg\"/>"

	stc := make(chan Tag)
	etc := make(chan Tag)

	go ParseChan(strings.NewReader(str), stc, etc)

	s := <-stc
	t.Log(s)
}
