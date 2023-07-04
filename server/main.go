package main

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	itemv1 "github.com/sofiastockholmcode/item-service/gen/item/v1"
	"github.com/sofiastockholmcode/item-service/gen/item/v1/itemv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	path, handler := itemv1connect.NewItemStoreServiceHandler(&itemStoreServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

// petStoreServiceServer implements the PetStoreService API.
type itemStoreServiceServer struct {
	itemv1connect.UnimplementedItemStoreServiceHandler
}

func (s *itemStoreServiceServer) GetItem(
	ctx context.Context,
	req *connect.Request[itemv1.GetItemRequest],
) (*connect.Response[itemv1.GetItemResponse], error) {
	id := req.Msg.GetItemId()
	log.Printf("Got a request to getItem with i %s", id)
	item := &itemv1.GetItemResponse{}
	item.Item = &itemv1.Item{
		ItemId: "1",
		Name:   "book",
	}
	return connect.NewResponse(item), nil
}
