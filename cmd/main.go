package main

import (
	"fmt"
	"os"

	cli_extension_lib_go "github.com/snyk/cli-extension-lib-go"
)

func main() {
	_, extensionInput, err := cli_extension_lib_go.InitExtension()
	if err != nil {
		fmt.Println("Error initializing extension")
		fmt.Println(err)
		os.Exit(1)
	}

	lang, err := extensionInput.Command.StringOptionValue("lang")
	if err != nil {
		fmt.Println("Invalid input")
		fmt.Println(err)

		os.Exit(1)
	}
	moreEmoji, err := extensionInput.Command.BoolOptionValue("more-emoji")
	if err != nil {
		fmt.Println("Invalid input")
		fmt.Println(err)
		os.Exit(1)
	}
	if moreEmoji {
		fmt.Println("😍")
		os.Exit(0)
	}
	fmt.Println("Woof in", lang)
	fmt.Println("Woof token:", extensionInput.Token)
}
