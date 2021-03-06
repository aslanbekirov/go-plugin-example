package main

import (
	"fmt"
	"os"
	"plugin"
)

type Greeter interface {
	Greet()
}

func main() {
	lang := "english"

	if len(os.Args) == 2 {
		lang = os.Args[1]
	}

	var mod string

	switch lang {
	case "english":
		mod = "./eng/eng.so"
	case "chienese":
		mod = "./chi/chi.so"
	default:
		fmt.Println("dont speak this language")
		os.Exit(1)
	}

	//load module
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//look for the symbol
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Check if loaded symbol is of desired type

	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)

	if !ok {
		fmt.Println("incorrect type of module type")
		os.Exit(1)
	}

	greeter.Greet()
}
