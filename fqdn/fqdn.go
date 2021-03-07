package fqdn

import (
	"fmt"
	"strings"
	"time"
)

const avaiable byte = 0

//Fqdn represents a fully qualified domain name avaiable status
type Fqdn struct {
	Status          byte      `json:"status"`
	Fqdn            string    `json:"fqdn"`
	AvaiableDomains []string  `json:"suggestions"`
	Reasons         []string  `json:"reasons"`
	ExpiresAt       time.Time `json:"expires-at"`
}

func (fqdn Fqdn) formatAvaiableDomains() string {
	var builder strings.Builder
	var domainName = strings.Split(fqdn.Fqdn, ".")

	for _, avaiableDomain := range fqdn.AvaiableDomains {
		builder.WriteString(fmt.Sprintf("\t- %s.%s\n", domainName[0], avaiableDomain))
	}

	return builder.String()
}

func (fqdn Fqdn) String() string {
	if fqdn.Status != avaiable {
		return fmt.Sprintf("Result: the domain %s is not avaiable\nReason: domain published\nExpiration: %v\nSuggestions:\n%s", fqdn.Fqdn, fqdn.ExpiresAt.Format("02/01/2006"), fqdn.formatAvaiableDomains())
	}

	return fmt.Sprintf("the domain %s is avaiable", fqdn.Fqdn)
}
