package main

import (
	"net/http"
	"net/http/httputil"

	log "github.com/Sirupsen/logrus"
)

type loggingRoundTripper struct {
	inner http.RoundTripper
}

func (d *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	d.dumpRequest(req)
	res, err := d.inner.RoundTrip(req)
	d.dumpResponse(res)
	return res, err
}

func (d *loggingRoundTripper) dumpRequest(r *http.Request) {
	dump, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		log.WithError(err).Error("oops")
	}
	log.WithField("section", "request").Infof(string(dump))
}

func (d *loggingRoundTripper) dumpResponse(r *http.Response) {
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.WithError(err).Error("oops")
	}
	log.WithField("section", "response").Infof(string(dump))
}
