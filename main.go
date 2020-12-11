package main

import "fmt"

func print(prg string) {
	fmt.Print(prg + "(\x60" + prg + "\x60)\n}\n")
}

func main() {
	print(`package main

import "fmt"

func print(prg string) {
	fmt.Print(prg + "(\x60" + prg + "\x60)\n}\n")
}

func main() {
	print`)
}
