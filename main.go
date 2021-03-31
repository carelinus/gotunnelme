package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/carelinus/gotunnelme/src/gotunnelme"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "gotunnelme <local address> <local port>")
		fmt.Fprintln(os.Stderr, "\nVersion:\n"+buildVersion(version, commit, date, builtBy))
		os.Exit(1)
	}
	address := os.Args[1]
	i, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	t := gotunnelme.NewTunnel()
	url, err := t.GetUrl("")
	if err != nil {
		panic(err)
	}
	print(url)
	err = t.CreateTunnel(address, i)
	if err != nil {
		panic(err)
	}
	t.StopTunnel()
}

func buildVersion(version, commit, date, builtBy string) string {
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
