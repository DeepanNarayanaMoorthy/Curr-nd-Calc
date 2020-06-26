package main

import (
    "bufio"
    "fmt"
    "flag"
    "net/http"
)

func main() {
    fromPtr := flag.String("from", "USD", "a string")
    toPtr := flag.String("to", "INR", "a string")
    storePtr := flag.String("store", "conv", "a string")
    flag.Parse()
    fmt.Println("from:", *fromPtr)
    fmt.Println("to:", *toPtr)
    fmt.Println("store in:", *storePtr)
    fmt.Println("Arguments List:", flag.Args())
    base_url:="https://api.exchangeratesapi.io/latest?base="
    resp, err := http.Get(base_url+*fromPtr)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }T

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
