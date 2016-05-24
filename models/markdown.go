package models

import ()

type Leaf interface {
}

type Node struct {
	Children []Node
}

type Page struct {
	Node
}

type Title struct {
	Node
}

type TextRange struct {
	Text string
}

func FindChild(node Title) {
}
