package client

import "fmt"

type Region string

const (
	RegionNA Region = "NA"
	RegionEU Region = "EU"
	RegionFE Region = "FE"
)

func Endpoint(region Region) (string, error) {
	switch region {
	case RegionNA:
		return "https://sellingpartnerapi-na.amazon.com", nil
	case RegionEU:
		return "https://sellingpartnerapi-eu.amazon.com", nil
	case RegionFE:
		return "https://sellingpartnerapi-fe.amazon.com", nil
	default:
		return "", fmt.Errorf("amazon sp-api: unsupported region %q", region)
	}
}
