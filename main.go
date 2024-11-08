package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func Custom(){
    fmt.Println("WebCheck")
	flag.PrintDefaults()
}

func flags() (string, string, string, bool, bool){
    address := flag.String("a", "", "Address of a host to check")
    addAddr := flag.String("n", "", "Add a host to monitor")
    delAddr := flag.String("d", "", "delete a host to monitor")
    listAddr := flag.Bool("l", false, "\n list of hosts to monitor")
    monitor := flag.Bool("m", false, "\nmonitor hosts")
    flag.Parse()
    return *address, *addAddr, *delAddr, *listAddr, *monitor
}

func monitorHost(addAddr bool){
    for{
    fmt.Print("\033[H\033[2J")

    file, err := os.Open("data.txt")
    if err != nil{
        fmt.Println(err)
    }
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        host := scanner.Text()
        resp, err := http.Get("http://" + host)
        if err != nil{
            fmt.Println(host, "- down")
        } else {
            fmt.Println(host, "- up")
            resp.Body.Close()
        }
    }
    time.Sleep(3 * time.Second)
    }
    

}

func addHost(addAddr string){
    file, _:= os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 6440)

    _, err := file.WriteString(addAddr + "\n")
    if err != nil{
        fmt.Println("error while adding to watchlist!")
        return
    }
    fmt.Println("added to watchlist!")
    defer file.Close()
}

func listHost(listAddr bool){
    var count int
    file, err := os.Open("data.txt")
    if err != nil{
        fmt.Println(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        count++
        fmt.Println(count, scanner.Text())
    }
    defer file.Close()
}

func delHost(deladdr string){

    // Otwórz plik do odczytu
    file, err := os.Open("data.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Tymczasowa lista przechowująca linie, które nie zawierają frazy
    var noweLinie []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linia := scanner.Text()
        if !strings.Contains(linia, deladdr) {
            noweLinie = append(noweLinie, linia)
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    plikZapis, err := os.Create("data.txt")
    if err != nil {
        panic(err)
    }
    defer plikZapis.Close()

    for _, linia := range noweLinie {
        _, err := plikZapis.WriteString(linia + "\n")
        if err != nil {
            panic(err)
        }
    }

    fmt.Println(deladdr+"has been deleted from a file")
}

func Check(address string){
    _, err := http.Get("http://" + address)
    if err != nil{
        fmt.Println(err)
    } else {
        fmt.Println("server is avaible!")
    }
}

func main(){
    _, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 6440)
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