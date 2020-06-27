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
    "os"
    "strconv"
    "strings"
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

type Finance struct {
    Amount int `json:"Amount"`
    Remark string `json:"Remark"`
    TransType string `json:"TransType"`
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
    finPtr := flag.String("fin", "", "Budget Manager")
    finresPtr := flag.String("finres", "n", "Returns amount of money left and budget report")
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
    } else if *finPtr !="" {
      resultt := strings.Split(*finPtr, ",")
      filenamee := "Finance.json"
      err := checkFile(filenamee)
      if err != nil {
          fmt.Println(err)    }

      file, err := ioutil.ReadFile(filenamee)
      if err != nil {
          fmt.Println(err)    }

      dataa := []Finance{}

      json.Unmarshal(file, &dataa)

      amt, _ := strconv.Atoi(resultt[0])
      newStruct := &Finance{
        Amount : amt,
        Remark : resultt[1],
        TransType : resultt[2],
      }

      dataa = append(dataa, *newStruct)

      dataBytes, err := json.Marshal(dataa)
      if err != nil {
          fmt.Println(err)    }

      err = ioutil.WriteFile(filenamee, dataBytes, 0644)
      if err != nil {
          fmt.Println(err)    }

    } else if *finresPtr == "y" {
      var arr []Finance
      finname := "Finance.json"
      //finresPtr := flag.String("finres", "n", "a string")
      //flag.Parse()
      datt, _ := ioutil.ReadFile(finname)
      json.Unmarshal([]byte(datt), &arr)
      amt:=0
      for index, element := range arr {
        fmt.Println(index, "=>", element)
        if element.TransType == "+" {
          amt=amt+element.Amount
        } else {
          amt=amt-element.Amount
        }

      }
      fmt.Println("Amount of money Left is : "+ strconv.Itoa(amt))

    }

}
