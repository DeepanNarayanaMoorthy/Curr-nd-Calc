package main

import (
    "bufio"
    "fmt"
    //"io/ioutil"
    "flag"
    "net/http"
    "encoding/json"
)
type Currency struct{
  Values map[string]float32
  Base string
  Datee string
}
func main() {
    var curr Currency
    fromPtr := flag.String("from", "USD", "From what currency?")
    toPtr := flag.String("to", "INR", "To what currency?")
    storePtr := flag.String("store", "conv", "Where should I store the value")
    flag.Parse()
    fmt.Println(*toPtr+*storePtr)
    base_url:="https://api.exchangeratesapi.io/latest?base="
    resp, err := http.Get(base_url+*fromPtr)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    fmt.Println("Response status:", resp.Status)
    scanner := bufio.NewScanner(resp.Body)
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    currJson:=""
    for scanner.Scan(){
        currJson=currJson+scanner.Text()
    }
    fmt.Println(currJson)
    json.Unmarshal([]byte(currJson), &curr)
    fmt.Println(curr.Base+" and "+curr.Datee)
    fmt.Println(curr.Values["CNY"])
}
