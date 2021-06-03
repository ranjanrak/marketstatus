package marketstatus

import (
        "encoding/json"
        "io/ioutil"
        "net/http"
        "fmt"
)

// MarketValue represents market time status for different segment
type MarketValue struct {
    MarketState []struct {
        Market              string  `json:"market"`
        MarketStatus        string  `json:"marketStatus"`
        TradeDate           string  `json:"tradeDate"`
        Index               string  `json:"index"`
        Last                float64 `json:"last"`
        Variation           float64 `json:"variation"`
        PercentChange       float64 `json:"percentChange"`
        MarketStatusMessage string  `json:"marketStatusMessage"`
    } `json:"marketState"`
}


func MarketStatus() MarketValue {

    var output MarketValue
    status_url := "https://www.nseindia.com/api/marketStatus"
    client := &http.Client{}
    req, err := http.NewRequest("GET", status_url, nil)
    if err != nil {
        fmt.Println("New request preparation failed", err)
    }
    req.Header.Set("User-Agent", 
            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Request failed", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    
    if err := json.Unmarshal(body, &output); err != nil {
        fmt.Println(err)
    }
    return output

}