// O(n), O(n)

// expected := "/foo/bar/baz"
// 	output := ShortenPath("/foo/../test/../test/../foo//bar/./baz")

// 1. check relative or absolute path (startwithSlash)
// 2. split into token by "/", so token can be foo, bar, . , .. , etc
// 3. check important token (eliminiate empty and .)
// 4. USE STACK to store update shortest path

package main

import (
	"fmt"
	"strings"
)

func ShortenPath(path string) string {
	// Write your code here.
	startsWithSlash := path[0] == '/'
	splits := strings.Split(path, "/")
	tokens := []string{}
	for _, token := range splits {
		if isImportantToken(token) {
			tokens = append(tokens, token)
		}
	}
	// fmt.Print(tokens)

	stack := []string{}
	if startsWithSlash {
		stack = append(stack, "")
	}
	for _, token := range tokens {
		if token == ".." {
			// if top is .. , just append, stack up ../../../
			// bcuz we keep .. for relative path
			if len(stack) == 0 || stack[len(stack)-1] == ".." {
				stack = append(stack, token)
				// if not root, remove top and go back to parent
			} else if stack[len(stack)-1] != "" {
				stack = stack[:len(stack)-1]
			}
			// always append token to stack
		} else {
			stack = append(stack, token)
		}
		fmt.Print(stack)
	}
	// edge case
	// wrong, bcuz if only / , return /
	// if len(stack) == 1 && len(stack) == 0 {
	if len(stack) == 1 && stack[0] == "" {
		return "/"
	}
	return strings.Join(stack, "/")
}

func isImportantToken(token string) bool {
	return len(token) > 0 && token != "."
}

func main() {
	input := "/foo/../test/../test/../foo//bar/./baz"
	fmt.Print(ShortenPath(input))
}
