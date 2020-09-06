package main

import (
	"context"
	"fmt"
	"grpc/proto"
	"grpc/store"
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type addservice struct{}

func (addservice) AddGarment(_ context.Context, in *proto.Item) (*proto.ItemList, error) {
	name := in.GetGarment()

	list := store.AddToList(name)

	fmt.Println("LIST FROM STORE", list)

	newList := proto.ItemList{}

	newList.List = []*proto.Item{}

	for _, item := range list {
		newItem := proto.Item{
			ID:      (int32)(item.ID),
			Garment: item.Garment,
		}
		newList.List = append(newList.List, &newItem)
	}

	fmt.Println(newList.GetList())
	return &proto.ItemList{List: newList.List}, nil
}
func (addservice) ClearList(_ context.Context, in *proto.ClearRequest) (*proto.ClearResponse, error) {
	fmt.Println("Clear List Called")
	store.Clear()

	return &proto.ClearResponse{}, nil

}

func (addservice) AllGarments(_ context.Context, in *proto.GetAllRequest) (*proto.ItemList, error) {

	list := store.GetAllItems()

	fmt.Println("Full List", list)

	newList := proto.ItemList{}

	newList.List = []*proto.Item{}

	for _, item := range list {
		newItem := proto.Item{
			ID:      (int32)(item.ID),
			Garment: item.Garment,
		}
		newList.List = append(newList.List, &newItem)
	}

	return &proto.ItemList{List: newList.List}, nil

}

func main() {
	server := grpc.NewServer()
	proto.RegisterGarmentServiceServer(server, addservice{})

	listener, err := net.Listen("tcp", ":5055")

	if err != nil {
		log.Fatal("Error while listening", err.Error())
	}

	go func() {
		log.Fatal("Serving gRPC:", server.Serve(listener).Error())
	}()

	grpcWebServer := grpcweb.WrapServer(server)
	httpServer := &http.Server{
		Addr: ":9003",
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")
				if grpcWebServer.IsGrpcWebRequest(r) {
					grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}

	log.Fatal("Serving Proxy", httpServer.ListenAndServe().Error())

}
