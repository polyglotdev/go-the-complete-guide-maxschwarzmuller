package main

import "fmt"

func main() {
	cloudProviders := map[string]string{
		"aws":     "Amazon Web Services",
		"gcp":     "Google Cloud Platform",
		"az":      "Microsoft Azure",
		"ibm":     "IBM Cloud",
		"alibaba": "Alibaba Cloud",
	}
	for key, value := range cloudProviders {
		fmt.Printf("%s:, %s\n", key, value)
	}
}
