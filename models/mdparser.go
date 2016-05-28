/**
Try to DIY refer to https://github.com/superhx/mark
so it will look quite similiar. :-(
**/
package models

import (
	"regexp"
)

type Parser struct {
	input      []byte
	blockRegex BlockRegex
	res_page   Page
}

type BlockRegex struct {
	newline, title_atx, title_setext, block_quotes *regexp.Regexp
}

func (p *Parser) Init() {
	p.blockRegex = BlockRegex{
		newline:      regexp.MustCompile(`^\n+`),
		title_atx:    regexp.MustCompile(`^(#{1,6}) *([^\n]+?) *#*(?:\n+|$)`),
		title_setext: regexp.MustCompile(`^ *([^\n]+)(?:\n+)([=-]+)(?:\n+|$)`),
		block_quotes: regexp.MustCompile(`^>([^\n]*)(?:\n|$)([^\n]+(?:\n+|$))`),
	}
}

func (p *Parser) Parse() {
	for len(p.input) > 0 {
		// TODO

	}
}
