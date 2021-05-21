package hello

import "rsc.io/quote/v3"

// Hello as a function
func Hello() string {
	return quote.HelloV3()
}

// Proverb concurrent as a function
func Proverb() string {
	return quote.Concurrency()
}
