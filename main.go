package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go-ipinfo/ipinfo"
)

func main() {

	ipAddress := flag.String("ip", "8.8.8.8", "Must use -ip or --ip")
	flag.Parse()

	authTransport := ipinfo.AuthTransport{Token: "2e63198b489866"}
	httpClient := authTransport.Client()
	client := ipinfo.NewClient(httpClient)
	info, err := client.GetInfo(net.ParseIP(*ipAddress))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
}
