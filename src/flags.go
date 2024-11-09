package main

import (
	"flag"
	"fmt"
)

func printUsage() {
	fmt.Println("WebCheck Usage:")
	flag.PrintDefaults()
}

func flags() (string, string, string, bool, bool) {
	address := flag.String("a", "", "Address of a host to check")
	addAddr := flag.String("n", "", "Add a host to monitor")
	delAddr := flag.String("d", "", "delete a host to monitor")
	listAddr := flag.Bool("l", false, "\n list of hosts to monitor")
	monitor := flag.Bool("m", false, "\nmonitor hosts")
	flag.Parse()
	return *address, *addAddr, *delAddr, *listAddr, *monitor
}
