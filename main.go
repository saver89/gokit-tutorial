package main

import (
	"gokit-tutorial/adapter"
	"gokit-tutorial/service"
	"gokit-tutorial/transports"
	"net/http"
	"os"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.StringService
	svc = adapter.StringService{}
	svc = loggingMiddleware{logger, svc}

	uppercaseHandler := httptransport.NewServer(
		transports.MakeUppercaseEndpoint(svc),
		transports.DecodeUppercaseRequest,
		transports.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		transports.MakeCountEndpoint(svc),
		transports.DecodeCountRequest,
		transports.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
}

type loggingMiddleware struct {
	logger log.Logger
	next   service.StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Count(s)
	return
}
