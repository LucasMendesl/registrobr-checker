package fqdn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//Search responsible to extract data from registro.br and check avaiability
func Search(host string) (*Fqdn, error) {
	host = url.PathEscape(host)
	res, err := http.Get(fmt.Sprintf("https://registro.br/v2/ajax/avail/raw/%s", host))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Ooops, request error with status %d", res.StatusCode)
	}

	var fqdn *Fqdn
	json.NewDecoder(res.Body).Decode(&fqdn)

	if len(fqdn.Reasons) > 0 {
		return nil, &fqdnError{
			fqdn:         fqdn.Fqdn,
			errorReasons: fqdn.Reasons,
		}
	}

	return fqdn, nil
}
