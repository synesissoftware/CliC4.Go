package main

import (
	clic4go "github.com/synesissoftware/CliC4.Go"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("CliC4.Go v%s\n", clic4go.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
