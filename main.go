package main

import (
	"fmt"
	"github.com/lttlrck/gotunnelme/src/gotunnelme"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "gotunnelme <local address> <local port>")
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
