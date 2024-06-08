// design browser system
package main

import "fmt"

func main() {
	browser := Constructor("leetcode.com")
	browser.Visit("google.com")
	fmt.Println(browser.Back(1)) // Output: "google.com"
	fmt.Println("facebook.com")
	fmt.Println(browser.Forward(1))

}

type BrowserHistory struct {
	stack   []string
	current int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		stack:   []string{homepage},
		current: 0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.stack = this.stack[:this.current+1]
	this.stack = append(this.stack, url)
	this.current++

}
func (this *BrowserHistory) Back(steps int) string {
	if steps > this.current {
		this.current = 0
	} else {
		this.current -= steps
	}
	return this.stack[this.current]

}

func (this *BrowserHistory) Forward(steps int) string {
	if steps > len(this.stack)-1-this.current {
		this.current = len(this.stack) - 1
	} else {
		this.current += steps
	}
	return this.stack[this.current]
}
