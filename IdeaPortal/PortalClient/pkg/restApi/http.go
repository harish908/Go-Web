package restApi

import (
	"PortalClient/pkg/tracing"
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	//log "github.com/sirupsen/logrus"
)

var client *http.Client

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 5 * time.Second,
		},
	}
}

// Ping sends a ping request to the given hostPort, ensuring a new span is created
// for the downstream call, and associating the span to the parent span, if available
// in the provided context.
func Ping(baseURL string, path string, apiMethod string, reqBody []byte, apiData chan []byte,
	apiErr chan error, ctx context.Context, operationName string) {
	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	url := baseURL + path
	req, err := http.NewRequest(apiMethod, url, bytes.NewReader(reqBody))
	//req.Header.Add("key", "value")
	if err != nil {
		apiData <- nil
		apiErr <- err
		return
	}

	if err := tracing.Inject(span, req); err != nil {
		apiData <- nil
		apiErr <- err
		return
	}

	// Send request and read data

	resp, err := client.Do(req)
	if err != nil {
		apiData <- nil
		apiErr <- err
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		apiData <- nil
		apiErr <- err
		return
	}

	//if resp.StatusCode != 200 {
	//	return "", fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	//}
	apiData <- body
	apiErr <- nil
}
