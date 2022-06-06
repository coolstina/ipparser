package main

import (
	"embed"
	"fmt"

	"github.com/coolstina/embedsfs"
	"github.com/coolstina/ipparser"
)

//go:embed embeds
var embeds embed.FS

func main() {
	embeds := embedsfs.NewEmbedsFS(embeds, "embeds")

	// For City Level IP Database
	city, err := ipparser.NewCity("embeds/path/to/17monipdb.datx", ipparser.WithEmbedsFS(embeds))
	if err == nil {
		fmt.Println(city.Find("8.8.8.8"))
		fmt.Println(city.Find("128.8.8.8"))
		fmt.Println(city.Find("255.255.255.255"))
		loc, err := city.FindLocation("27.190.252.103")
		if err == nil {
			fmt.Println(string(loc.ToJSON()))
			// Output:
			/*
			   {
			       "Country": "China",
			       "Province": "Hebei",
			       "City": "Tangshan",
			       "Organization": "",
			       "ISP": "ChinaTelecom",
			       "Latitude": "39.635113",
			       "Longitude": "118.175393",
			       "TimeZone": "Asia/Shanghai",
			       "TimeZone2": "UTC+8",
			       "CityCode": "130200",
			       "PhonePrefix": "86",
			       "CountryCode": "CN",
			       "ContinentCode": "AP",
			       "IDC": "",
			       "BaseStation": "",
			       "Anycast": false
			   }
			*/
		}
	}

	// Only China District IP Database
	dis, err := ipparser.NewDistrict("embeds/path/to/quxian.datx", ipparser.WithEmbedsFS(embeds))
	if err == nil {
		fmt.Println(dis.Find("1.12.46.0"))
		fmt.Println(dis.Find("223.255.127.0"))
	}

	// Only China Base Station IP Database
	bst, err := ipparser.NewBaseStation("embeds/path/to/station_ip.datx", ipparser.WithEmbedsFS(embeds))
	if err == nil {
		fmt.Println(bst.Find("1.30.6.0"))
		fmt.Println(bst.Find("223.221.121.0"))
		fmt.Println(bst.Find("223.221.121.255"))
	}
}
