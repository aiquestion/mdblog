/**
Try to DIY refer to https://github.com/superhx/mark
so it will look quite similiar. :-(
**/
package models

import (
	"regexp"
	"strings"
)

type Parser struct {
	input      []byte
	blockRegex BlockRegex
	res_page   Page
}

type BlockRegex struct {
	newline, title_atx, title_setext, block_quotes *regexp.Regexp
	list, indent_block                             *regexp.Regexp
}

func (p *Parser) Init() {
	no_emptyline := `[^\n]+(?:\n|$)`
	line_startwith_indent := `[ {4}\t]([])`
	line_startwith_indent = strings.Replace(line_startwith_indent, "no_emptyline", no_emptyline, -1)
	p.blockRegex = BlockRegex{
		newline:      regexp.MustCompile(`^\n+`),
		title_atx:    regexp.MustCompile(`^(#{1,6}) *([^\n]+?) *#*(?:\n+|$)`),
		title_setext: regexp.MustCompile(`^ *([^\n]+)(?:\n+)([=-]+)(?:\n+|$)`),
		block_quotes: regexp.MustCompile(`^>([^\n]*)(?:\n|$)`),
		list:         regexp.MustCompile(`^(\*|\+|\-|(?:\d+\.)) *([^\n]+)(?:\n+|$)`),
		indent_block: regexp.MustCompile(`^(?: {4}|\t)([^\n]+)(?:\n|$)`),
	}
}

func (p *Parser) Parse() {
	for len(p.input) > 0 {
		if m := p.blockRegex.newline.Find(p.input); m != nil {
			p.input = p.input[len(m):]
			continue
		}

		if m := p.blockRegex.title_atx.FindSubmatch(p.input); m != nil {

			continue
		}
	}
}

func (p *Parser) handleTitleAtx(ms [][]byte) {
}
