package gateways

import (
	"PortalClient/pkg/restApi"
	"context"

	"github.com/spf13/viper"
)

func GetIdeas(ctx context.Context) ([]byte, error) {
	apiData := make(chan []byte)
	apiErr := make(chan error)
	baseURL := viper.GetString("PortalServer.baseURL")
	go restApi.Ping(baseURL, "ideas", "GET", nil, apiData, apiErr, ctx, "getIdeas")
	//log.Info("Line executes before ping function, since we used go routines")
	return <-apiData, <-apiErr
}

func PostIdea(body []byte, ctx context.Context) ([]byte, error) {
	apiData := make(chan []byte)
	apiErr := make(chan error)
	baseURL := viper.GetString("PortalServer.baseURL")
	go restApi.Ping(baseURL, "postIdea", "POST", body, apiData, apiErr, ctx, "postIdea")
	//log.Info("Line executes before ping function, since we used go routines")
	return <-apiData, <-apiErr
}
