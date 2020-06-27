package main

import (
    "bufio"
    "fmt"
    "flag"
    "net/http"
    "encoding/json"
    "go/token"
  	"go/types"
    "io/ioutil"
    "strings"
    "os"
)
type Currency struct{
  Rates map[string]float32
  Base string `json:"base"`
  Date string `json:"date"`
}

type Contactt struct {
    Name string `json:"Name"`
    Number string `json:"Number"`
    Relation string `json:"Relation"`
}

func checkFile(filename string) error {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        _, err := os.Create(filename)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    var curr Currency
    fromPtr := flag.String("from", "", "From what currency?")
    toPtr := flag.String("to", "", "To what currency?")
    calcPtr := flag.String("calc", "", "Simple Calculator")
    contPtr := flag.String("cont", "", "Simple Contact Manager")
    seecontPtr := flag.String("seecont", "", "see contents of Simple Contact Manager")
    flag.Parse()
    if *calcPtr != "" {
      fs := token.NewFileSet()
  		tv, _ := types.Eval(fs, nil, token.NoPos, *calcPtr)
  		fmt.Println(tv.Value)
    } else if *fromPtr != "" {
      fmt.Println(*toPtr)
      base_url:="https://api.exchangeratesapi.io/latest?base="
      resp, err := http.Get(base_url+*fromPtr)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()
      scanner := bufio.NewScanner(resp.Body)
      if err := scanner.Err(); err != nil {
          panic(err)
      }
      currJson:=""
      for scanner.Scan(){
          currJson=currJson+scanner.Text()
      }
      json.Unmarshal([]byte(currJson), &curr)
      if *toPtr == "all" {
        fmt.Println(curr.Rates)
      } else {
        fmt.Println(curr.Rates[*toPtr])
      }
    } else if *contPtr != "" {
      result := strings.Split(*contPtr, ",")
      filename := "Contacts.json"
      err := checkFile(filename)
      if err != nil {
          fmt.Println(err)
      }

      file, err := ioutil.ReadFile(filename)
      if err != nil {
          fmt.Println(err)
      }

      data := []Contactt{}

      json.Unmarshal(file, &data)

      newStruct := &Contactt{
          Name: result[0],
          Number: result[1],
          Relation: result[2],
      }

      data = append(data, *newStruct)

      dataBytes, err := json.Marshal(data)
      if err != nil {
          fmt.Println(err)
      }

      err = ioutil.WriteFile(filename, dataBytes, 0644)
      if err != nil {
          fmt.Println(err)
      }

    } else if *seecontPtr == "y" {
      dat, _ := ioutil.ReadFile("Contacts.json")
      fmt.Print(string(dat))
    }

}
