package fqdn

import (
	"fmt"
	"strings"
)

type fqdnError struct {
	fqdn         string
	errorReasons []string
}

func (err fqdnError) Error() string {
	return fmt.Sprintf("fqdn [%s] terminated with error(s): %s\n", err.fqdn, strings.Join(err.errorReasons, "\n"))
}
