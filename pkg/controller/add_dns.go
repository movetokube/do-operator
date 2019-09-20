package controller

import (
	"github.com/movetokube/do-operator/pkg/controller/dns"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, dns.Add)
}
