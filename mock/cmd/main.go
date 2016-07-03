package main

import (
	"flag"
	"fmt"

	"github.com/suzuken/misc/mock"
)

func main() {
	var (
		user = flag.String("user", "suzuken", "gist user name")
	)
	flag.Parse()
	urls, err := mock.ListGists(*user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", urls)
}
