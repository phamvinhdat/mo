package main

import (
	"fmt"
	"os"

	"github.com/phamvinhdat/heroyaml"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("src is required")
	}

	res, err := heroyaml.K8sEnvConverter{}.Convert(args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
