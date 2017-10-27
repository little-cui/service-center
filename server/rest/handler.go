package rest

import (
	"github.com/ServiceComb/service-center/pkg/rest"
	roa "github.com/ServiceComb/service-center/pkg/rest"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

var (
	defaultHandler ServerHandler
)

func init() {
	defaultHandler.ROAServerHandler = roa.GetRouter().(*roa.ROAServerHandler)

	// api
	http.Handle("/", &defaultHandler)

	// prometheus metrics
	http.Handle("/metrics", prometheus.Handler())
}

type ServerHandler struct {
	*rest.ROAServerHandler
}
