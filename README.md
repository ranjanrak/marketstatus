# marketstatus
A utility to know current market open or close status along with next trade date and index info.

## Installation
```
go get -u github.com/ranjanrak/marketstatus
```

## Usage
```
package main

import (
    "fmt"
    "encoding/json"
    "github.com/ranjanrak/marketstatus"
)

func main() {
    status := marketstatus.MarketStatus()
    fmt.Printf("%+v\n", status)
    //Pretty print by marshalling the structure
    jsonF, _ := json.Marshal(message)
    fmt.Println(string(jsonF))
}
```

## Response
```
{MarketState:[{Market:Capital Market MarketStatus:Open TradeDate:04-Jun-2021 
Index:NIFTY 50 Last:15680.4 Variation:-9.950000000000728 PercentChange:-0.06 
MarketStatusMessage:Normal Market is Open} 
{Market:Currency MarketStatus:Open TradeDate:04-Jun-2021 
Index: Last:0 Variation:0 PercentChange:0 MarketStatusMessage:Market is Open} 
{Market:Commodity MarketStatus:Open TradeDate:04-Jun-2021 
Index: Last:0 Variation:0 PercentChange:0 MarketStatusMessage:Market is Open} 
{Market:Debt MarketStatus:Open TradeDate:04-Jun-2021 
Index: Last:0 Variation:0 PercentChange:0 MarketStatusMessage:Market is Open}]}
```
Pretty print by marshalling
```
{"marketState":[{"market":"Capital Market","marketStatus":"Open","tradeDate":"04-Jun-2021",
"index":"NIFTY 50","last":15648.3,"variation":-42.05000000000109,"percentChange":-0.27,
"marketStatusMessage":"Normal Market is Open"},{"market":"Currency","marketStatus":"Open",
"tradeDate":"04-Jun-2021","index":"","last":0,"variation":0,"percentChange":0,
"marketStatusMessage":"Market is Open"},
{"market":"Commodity","marketStatus":"Open","tradeDate":"04-Jun-2021","index":"","last":0,
"variation":0,"percentChange":0,"marketStatusMessage":"Market is Open"},
{"market":"Debt","marketStatus":"Open","tradeDate":"04-Jun-2021","index":"","last":0,
"variation":0,"percentChange":0,"marketStatusMessage":"Market is Open"}]}
```