package main

import (
	"github.com/lapingvino/etime"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Call("alert", etime.Now().String() + " " + etime.Now().Element().String())
}
