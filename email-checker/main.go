package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scnner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, has MX, SPF,sprRecord, DMARC, dmarcRecord/n")
	for scnner.Scan() {
		checkDomain(scnner.Text())
	}
	if err := scnner.Err(); err != nil {
		log.Fatal("Error: could not read form input %v", err)
	}

}
func checkDomain(domain string) {
	var hasMX, SPF, DMARC bool
	var spf, dmarc string
	mxs, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("ERROR: %v/n", err)
	}
	if len(mxs) > 0 {
		hasMX = true
	}
	txtr, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v/n", err)
	}
	for _, v := range txtr {
		if strings.HasPrefix(v, "v=spf1") {
			SPF = true
			spf = v
			break
		}
	}
	dmarcs, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error %v /n", err)
	}
	for _, v := range dmarcs {
		if strings.HasPrefix(v, "v=DMARC1") {
			DMARC = true
			dmarc = v
			break
		}
	}
}
