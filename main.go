package main

import "flag"

func main() {
	search := flag.String("s", "", "specify ip range to discover")
	flag.Parse()

}
