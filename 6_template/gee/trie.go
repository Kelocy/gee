package gee

import "strings"

type node struct {
	pattern  string  // /p/:lang
	part     string  // :lang
	children []*node // sub node [doc, tutorial, intro]
	isWild   bool    // exact match, is True when part include : or *
}

// First successfully matched node, for insert
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || ((part[0] == ':' || part == "") && child.isWild) {
			return child
		}
	}
	return nil
}

// All successfully matched nodes, for search
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	wildNodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part {
			nodes = append(nodes, child)
		} else if child.isWild {
			wildNodes = append(wildNodes, child)
		}
	}
	nodes = append(nodes, wildNodes...)
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		// marks this is a true pattern
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	// need to create a new node
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// just a path, not our marked node
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
