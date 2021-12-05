package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"Posting-Feed-be/datastruct"
	"Posting-Feed-be/logging"
	"Posting-Feed-be/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type AphService interface {
	//HelloWorldService(context.Context, string) string
	FeedCreateService(context.Context, string, string) string
}

type aphService struct{}

var ErrEmpty = errors.New("empty string")

func (aphService) FeedCreateService(_ context.Context, image_feed string, caption_feed string) string {

	return call_ServiceCreateFeed(image_feed, caption_feed)
}

func call_ServiceCreateFeed(image_feed string, caption_feed string) string {

	messageResponse := service.Create(image_feed, caption_feed)

	return messageResponse

}

func makeFeedCreateEndpoint(aph AphService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(datastruct.FeedCreateRequest)
		logging.Log(fmt.Sprintf("Link %s, Caption %s", req.IMAGE_FEED, req.CAPTION_FEED))
		v := aph.FeedCreateService(ctx, req.IMAGE_FEED, req.CAPTION_FEED)
		logging.Log(fmt.Sprintf("Response Final Message %s", v))
		return datastruct.FeedCreateResponse{200, v}, nil
	}
}

func decodeFeedCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request datastruct.FeedCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHttpsServicesAndStartListener() {
	aph := aphService{}

	FeedCreateHandler := httptransport.NewServer(
		makeFeedCreateEndpoint(aph),
		decodeFeedCreateRequest,
		encodeResponse,
	)

	http.Handle("/feed/create", FeedCreateHandler)
}
