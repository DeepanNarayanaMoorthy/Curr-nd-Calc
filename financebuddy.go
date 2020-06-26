package main

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {
    base_url:="https://api.exchangeratesapi.io/latest?base="
    usd:="USD"
    stringg:=base_url+usd
    fmt.Println(stringg)
    resp, err := http.Get(base_url+usd)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
