package main

import (
	"strings"

	"github.com/rossedman/shellama/cmd"
)

var (
	commit  string
	version string
)

func main() {
	cmd.Execute(strings.Join([]string{version, commit}, "+"))
}
