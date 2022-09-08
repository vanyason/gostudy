package algo

import "golang.org/x/net/html"

func DfsRecursive(n *html.Node, f func(n *html.Node)) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		f(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		DfsRecursive(c, f)
	}
}

func DfsStackbased(n *html.Node, f func(n *html.Node)) {
	var stack []*html.Node
	stack = append(stack, n)

	for len(stack) != 0 {
		iLast := len(stack) - 1
		n = stack[iLast]
		stack = stack[:iLast]

		if n == nil {
			continue
		}

		if n.Type == html.ElementNode {
			f(n)
		}

		for c := n.LastChild; c != nil; c = c.PrevSibling {
			stack = append(stack, c)
		}
	}
}
