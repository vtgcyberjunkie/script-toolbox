package main

import (
    "context"
    "os"
    "fmt"
    "github.com/digitalocean/godo"
    "io/ioutil"
    "net/http"
)


func pub_ip() string {  
        // https://gist.github.com/ankanch/8c8ec5aaf374039504946e7e2b2cdf7f
        url := "https://api.ipify.org?format=text"      // we are using a pulib IP API, we're using ipify here, below are some others
                                              // https://www.ipify.org
                                              // http://myexternalip.com
                                              // http://api.ident.me
                                              // http://whatismyipaddress.com/api
        resp, err := http.Get(url)
        if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
        ip, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                panic(err)
        }
        return string(ip)
}

func main() {
    var my_host = ""
    var my_domain = ""
    var my_domain_id=285678041
    
    if len(os.Args) >= 3 { 
      my_host = os.Args[1]
      my_domain = os.Args[2]
    } else {
      fmt.Println("Usage: ./script <hostname> <domain>")
      os.Exit(-1)
    }
  
    token := os.Getenv("DIGITALOCEAN_TOKEN")

    client := godo.NewFromToken(token)
    ctx := context.TODO()

    editRequest := &godo.DomainRecordEditRequest{
      Type: "A",
      Name: my_host,
      Data: pub_ip(),
    }

    domainRecord, _, err := client.Domains.EditRecord(ctx, my_domain, my_domain_id, editRequest)
    fmt.Println(domainRecord,err)
}
