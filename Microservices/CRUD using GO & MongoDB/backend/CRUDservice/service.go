package main

import (
	"context"
	"errors"
	"fmt"
	"go_blog/crud_func/global"
	"go_blog/crud_func/proto"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type crudServer struct{}

var server crudServer
var detailsFiltered []bson.M
var nname string
var user global.User

func (crudServer) CreateEmp(_ context.Context, in *proto.CreateEmpRequest) (*proto.CreateEmpResponse, error) {
	id, name, level, stream := in.GetId(), in.GetName(), in.GetLevel(), in.GetStream()
	newUser := global.User{ID: primitive.NewObjectID(), EID: id, Empname: name, Level: level, Stream: stream} //creates new user
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	_, err := global.DB.Collection("user").InsertOne(ctx, newUser) //inserts one user
	if err != nil {
		log.Println("Error Inserting user :", err.Error())
		return &proto.CreateEmpResponse{}, errors.New("Something didn't go as planned")
	}
	return &proto.CreateEmpResponse{Token: newUser.GetToken()}, nil
}

func (crudServer) ReadEmp(_ context.Context, in *proto.ReadEmpRequest) (*proto.ReadEmpResponse, error) {
	id := in.GetId()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	global.DB.Collection("user").FindOne(ctx, bson.M{"eid": id}).Decode(&user)

	return &proto.ReadEmpResponse{Token: user.GetToken()}, nil
}

func (crudServer) UpdateEmp(_ context.Context, in *proto.UpdateEmpRequest) (*proto.UpdateEmpResponse, error) {
	id := in.GetId()
	nname := in.GetNname()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	res, err := global.DB.Collection("user").UpdateOne( //Update One document
		ctx,
		bson.M{"eid": id},
		bson.D{
			{"$set", bson.D{{"ename", nname}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents! \n", res.ModifiedCount)
	return &proto.UpdateEmpResponse{Token: user.GetToken()}, nil
}

func (crudServer) DeleteEmp(_ context.Context, in *proto.DeleteEmpRequest) (*proto.DeleteEmpResponse, error) {
	id := in.GetId()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	filter := bson.M{"eid": id}

	res, err := global.DB.Collection("user").DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v Documents! \n", res.DeletedCount)
	return &proto.DeleteEmpResponse{Token: user.GetToken()}, nil
}

func (crudServer) AuthUser(_ context.Context, in *proto.AuthUserRequest) (*proto.AuthUserResponse, error) {
	token := in.GetToken()
	user := global.UserFromToken(token)
	return &proto.AuthUserResponse{Eid: user.EID, Empname: user.Empname, Level: user.Level, Stream: user.Stream}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterEmpServiceServer(server, crudServer{})
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal("error Creating listener", err.Error())
	}
	go func() {
		log.Fatal("Serving gRPC: ", server.Serve(listener).Error())
	}()
	grpcWebServer := grpcweb.WrapServer(server)

	httpServer := &http.Server{
		Addr: ":9001",
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
	log.Fatal("Serving Proxy: ", httpServer.ListenAndServe().Error())
}

//Run: go get golang.org/x/net/http2/h2c
//Run: go get github.com/improbable-eng/grpc-web/go/grpcweb
