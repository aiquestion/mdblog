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
	newline, title_atx *regexp.Regexp
}

func (p *Parser) Init() {
	p.blockRegex = BlockRegex{
		newline:   regexp.MustCompile(`^\n+`),
// TODO some error here
		title_atx: regexp.MustCompile(`^#{1,6}([^\n]*)#*\n`),
	}
}

func (p *Parser) Parse() {
	for len(p.input) > 0 {
		// TODO

	}
}
