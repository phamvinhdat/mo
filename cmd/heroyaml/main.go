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

	res, err := heroyaml.K8sEnvConverter{}.Convert(`
# All configuration options of the service will be specified here.

# Service's info.
service:
  cluster: local
  namespace: default
  tenants:
    - default
    - portfolio
    - datpv
`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
