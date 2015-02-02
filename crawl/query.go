package crawl

import (
	"fmt"
	"net/http"
	"time"
)

type Query struct {
	From     string // The page where the link was found
	Request  *http.Request
	Response *http.Response
	Latency  time.Duration
	// TODO complete HTML - line number?
}

func (q Query) String() string {
	return fmt.Sprintf(
		"%d (%.3f ms): %s FROM %s",
		q.Response.StatusCode,
		float64(q.Latency)/1e6,
		q.Request.URL,
		q.From,
	)
}
