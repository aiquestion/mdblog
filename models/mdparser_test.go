package models

import (
	"bytes"
	"testing"
)

func assert(a, b []byte, t *testing.T) {
	if bytes.Compare(a, b) != 0 {
		t.Error(string(a), string(b), "doesn't matched")
	}
}

func TestNewLineRegex(t *testing.T) {
	s := []byte("\n\n\n")
	p := new(Parser)
	p.Init()
	m := p.blockRegex.newline.Find(s)
	t.Log("matched string:", m)
	if bytes.Compare(s, m) != 0 {
		t.Error("newline unmatched")
	}
}

func TestTitleAtxRegex(t *testing.T) {
	s := []byte("###  i'm a title  ######\n\n")
	p := new(Parser)
	p.Init()
	m := p.blockRegex.title_atx.FindSubmatch(s)
	t.Log("matched string:", m)
	assert([]byte("###"), m[1], t)
	assert([]byte("i'm a title"), m[2], t)
}

func TestTitleSetextRegex(t *testing.T) {
	s := []byte("  i'm a title\n\n===========\n\n")
	p := new(Parser)
	p.Init()
	m := p.blockRegex.title_setext.FindSubmatch(s)
	t.Log("matched string:", m)
	assert([]byte("i'm a title"), m[1], t)
	assert([]byte("==========="), m[2], t)
}

func TestBlockQuotesRegex(t *testing.T) {
	s := []byte("> this is a quotes\n")
	p := new(Parser)
	p.Init()
	m := p.blockRegex.block_quotes.FindSubmatch(s)
	t.Log("matched string:", m)
	assert([]byte(" this is a quotes"), m[1], t)
}

func TestListNumberRegex(t *testing.T) {
	s := []byte("90. this is a list\n")
	p := new(Parser)
	p.Init()
	m := p.blockRegex.list.FindSubmatch(s)
	t.Log("matched string:", m)
	assert([]byte("90."), m[1], t)
	assert([]byte("this is a list"), m[2], t)
}

func TestListRegex(t *testing.T) {
	s := []byte("*   this is a list\n")
	p := new(Parser)
	p.Init()
	m := p.blockRegex.list.FindSubmatch(s)
	t.Log("matched string:", m)
	assert([]byte("*"), m[1], t)
	assert([]byte("this is a list"), m[2], t)
}
