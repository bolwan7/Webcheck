package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func monitorHost(addAddr bool) {
	dataFilePath, _ := getDataFilePath()

	for {
		fmt.Print("\033[H\033[2J")

		file, err := os.Open(dataFilePath)
		if err != nil {
			fmt.Println(err)
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			host := scanner.Text()
			resp, err := http.Get("http://" + host)
			if err != nil {
				fmt.Println(host, "- down")
			} else {
				fmt.Println(host, "- up")
				resp.Body.Close()
			}
		}
		time.Sleep(3 * time.Second)
	}

}

