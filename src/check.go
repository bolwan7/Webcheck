package main
import (
	"fmt"
	"net/http"	
)

func Check(address string) {
	_, err := http.Get("http://" + address)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("server is avaible!")
	}
}
