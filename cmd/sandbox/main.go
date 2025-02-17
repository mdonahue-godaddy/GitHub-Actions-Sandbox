package main

import (
	"fmt"

	"github.com/mdonahue-godaddy/GitHub-Actions-Sandbox/version"
)

func main() {
	// ctx := context.Background()

	vrsn := version.GetLongVersion()

	fmt.Println(vrsn)
}
