package main

import "os"

func main() {
	if err := os.MkdirAll("models", os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("results", os.ModePerm); err != nil {
		panic(err)
	}
	Values()
	Social()
}
