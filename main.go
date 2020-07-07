package main

import (
	"syscall/js"
)

func print(this js.Value, i []js.Value) interface{} {
	return "return value"
}

func registerCallbacks() {
	js.Global().Set("print", js.FuncOf(print))
}

var currentText = "Before"

func changeText(this js.Value, i []js.Value) interface{} {
	if currentText == "Before" {
		currentText = "After"
	} else {
		currentText = "Before"
	}
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", currentText)
	return nil
}

func changeTextCallbacks() {
	js.Global().Set("changeText", js.FuncOf(changeText))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	changeTextCallbacks()
	<-c
}
