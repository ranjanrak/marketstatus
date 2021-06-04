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

    "github.com/ranjanrak/marketstatus"
)

func main() {
    status := marketstatus.MarketStatus()
    fmt.Printf("%+v\n", status)
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