package router

import "strings"

type node struct {
	url      string  //only leaf node has url
	part     string  //current node's part of url
	children []*node //child nodes
	isWild   bool    //is ':' or '*'
}

func (n *node) getChild(part string) *node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

func (n *node) getAllChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(url string, parts []string, height int) {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		n.url = url
		return
	}

	// current part
	part := parts[height]
	// try to get an eligible child
	child := n.getChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(url, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		return n
	}

	part := parts[height]
	children := n.getAllChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
