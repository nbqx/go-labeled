package main

import (
	"fmt"
	"os"

	"github.com/nbqx/go-labeled/label"
)

func main() {
	var arg string
	label := label.NewLabel()

	if len(os.Args) == 2 {
		arg = os.Args[1]
	}
	
	if res,err := label.Get(arg); err == nil {
		fmt.Printf("%s",res)
	}else{
		fmt.Errorf("%s",err)
	}
}
