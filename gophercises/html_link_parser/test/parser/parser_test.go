package parser

import (
	"strings"
	"testing"

	"htmllinkparser/internal/examples"
	"htmllinkparser/internal/parser"
)

func TestParserEx1(t *testing.T) {
	linksActual, err := parser.Parse(strings.NewReader(examples.Ex1))
	if err != nil {
		t.Error("can not parse: ", err)
	}

	linksExpected := []parser.Link{
		{
			Href: "/other-page",
			Text: "A link to a page",
		},
		{
			Href: "/page-two",
			Text: "A link to another page",
		},
	}

	if len(linksActual) != len(linksExpected) {
		t.Fatalf("different length.\nexpected (%d): %+v\nactual (%d): %+v", len(linksExpected), linksExpected, len(linksActual), linksActual)
	}

	for i, _ := range linksActual {
		if linksActual[i] != linksExpected[i] {
			t.Errorf("links mismatch.\nexpected: %+v\nactual: %+v", linksExpected[i], linksActual[i])
		}
	}
}

func TestParserEx2(t *testing.T) {
	linksActual, err := parser.Parse(strings.NewReader(examples.Ex2))
	if err != nil {
		t.Error("can not parse: ", err)
	}

	linksExpected := []parser.Link{
		{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
		{
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on Github!",
		},
	}

	if len(linksActual) != len(linksExpected) {
		t.Fatalf("different length.\nexpected (%d): %+v\nactual (%d): %+v", len(linksExpected), linksExpected, len(linksActual), linksActual)
	}

	for i, _ := range linksActual {
		if linksActual[i] != linksExpected[i] {
			t.Errorf("links mismatch.\nexpected: %+v\nactual: %+v", linksExpected[i], linksActual[i])
		}
	}
}

func TestParserEx3(t *testing.T) {
	linksActual, err := parser.Parse(strings.NewReader(examples.Ex3))
	if err != nil {
		t.Error("can not parse: ", err)
	}

	linksExpected := []parser.Link{
		{
			Href: "#",
			Text: "Login",
		},
		{
			Href: "/lost",
			Text: "Lost? Need help?",
		},
		{
			Href: "https://twitter.com/marcusolsson",
			Text: "@marcusolsson",
		},
	}

	if len(linksActual) != len(linksExpected) {
		t.Fatalf("different length.\nexpected (%d): %+v\nactual (%d): %+v", len(linksExpected), linksExpected, len(linksActual), linksActual)
	}

	for i, _ := range linksActual {
		if linksActual[i] != linksExpected[i] {
			t.Errorf("links mismatch.\nexpected: %+v\nactual: %+v", linksExpected[i], linksActual[i])
		}
	}
}

func TestParserEx4(t *testing.T) {
	linksActual, err := parser.Parse(strings.NewReader(examples.Ex4))
	if err != nil {
		t.Error("can not parse: ", err)
	}

	linksExpected := []parser.Link{
		{
			Href: "/dog-cat",
			Text: "dog cat",
		},
	}

	if len(linksActual) != len(linksExpected) {
		t.Fatalf("different length.\nexpected (%d): %+v\nactual (%d): %+v", len(linksExpected), linksExpected, len(linksActual), linksActual)
	}

	for i, _ := range linksActual {
		if linksActual[i] != linksExpected[i] {
			t.Errorf("links mismatch.\nexpected: %+v\nactual: %+v", linksExpected[i], linksActual[i])
		}
	}
}
