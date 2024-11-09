package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


func addHost(addAddr string){
	dataFilePath, _ := getDataFilePath()

    file, _ := os.OpenFile(dataFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 6440)

    _, err := file.WriteString(addAddr + "\n")
    if err != nil{
        fmt.Println("error while adding to watchlist!")
        return
    }
    fmt.Println("added to watchlist!")
    defer file.Close()
}

func listHost(listAddr bool){
	dataFilePath, _ := getDataFilePath()

    var count int
    file, err := os.Open(dataFilePath)
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