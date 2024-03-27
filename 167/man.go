/*

Problem Statement:
Given an absolute path for a file (Unix-style), simplify it. Or in other words, 
convert it to the canonical path.

In a Unix-style file system, a period '.' refers to the current directory. Furthermore, 
a double period '..' moves the directory up a level. For more information, 
see: Absolute path vs relative path in Linux/Unix

Note that the returned canonical path must always begin with a slash '/', 
and there must be only a single slash '/' between two directory names. 
The last directory name (if it exists) must not end with a trailing '/'. Also,
 the canonical path must be the shortest string representing the absolute path.

Example:
Input: "/a/./b/../../c/"
Output: "/c"
Explanation:
Starting from the root directory, we move to 'a', then to '.' (current directory), then to 'b', then to '..' (parent directory), and finally to 'c' which is the canonical path.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/home/"
	fmt.Printf("result:%v\n", simplifyPath(path))

}

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	stack := []string{}

	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		} else if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, dir)
		}
	}

	return "/" + strings.Join(stack, "/")
}
