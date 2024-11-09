package main

import (
	"flag"
	"fmt"
	"os"
)

func Custom(){
    fmt.Println("WebCheck")
	flag.PrintDefaults()
}







func main(){
    dataFilePath, _ := getDataFilePath()

    _, err := os.OpenFile(dataFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 6440)
    if err != nil{
        fmt.Println(err)
    }

    Custom()

    address, addAddr, delAddr, listAddr, monitor:= flags()
    if address != "" {
        Check(address)
    }

    if addAddr != "" {
        addHost(addAddr)
    }
    if delAddr != ""{
        delHost(delAddr)
    }
    if listAddr == true{
        listHost(listAddr)
    }

    if monitor == true{
        monitorHost(monitor)
    }
    


}