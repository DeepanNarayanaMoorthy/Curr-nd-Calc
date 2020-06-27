# curr-n-calc

## This is a Command Line Based Currency Converter and calculator written in GoLang

 :black_circle: ### For a particular Currency

```
go build curr-nd-calc.go
./curr-nd-calc -from=USD -to=INR 
```

 :black_circle: ### For All Currency

```
go build curr-nd-calc.go
./curr-nd-calc -from=USD -to=all 
```

 :black_circle: ### For Calculator

```
go build curr-nd-calc.go
./curr-nd-calc -calc=3*(1+2(3*4)-22) 
```

 :black_circle: ### For Adding Contact

```
go build curr-nd-calc.go
./curr-nd-calc -cont=JohnDoe,9876543210,MyDad 
```

 :black_circle: ### For Adding Budget info
#### Lets say that I have spent $234 in laptop service, This is expenditure so " - " (minus) sign will be specified
```
go build curr-nd-calc.go
./curr-nd-calc -fin=234,laptop service,- 
```
#### If I have got $234 as loan from  a friend, This is return so " + " (plus) sign will be specified
```
go build curr-nd-calc.go
./curr-nd-calc -fin=234,loanfromfriend,+ 
```

 :black_circle: ### For calculating profit or loss and see budget info 

```
go build curr-nd-calc.go
./curr-nd-calc -finres=y
```

#### PS : This is just a small code, feel free to comment or contribute  :wink:
