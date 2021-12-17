package main

import (
	"context"
	"errors"
	"fmt"
	"go_blog/crudfunc/global"
	"go_blog/crudfunc/proto"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type crudServer struct{}

func (server crudServer) Create(_ context.Context, in *proto.CreateRequest) (*proto.AuthResponse, error) {
	eid, empname, elevel, estream := in.GetEid(), in.GetEmpname(), in.GetElevel(), in.GetEstream()
	if len(eid) != 8 || len(estream) > 20 || elevel > 13 || len(empname) < 8 {
		return &proto.AuthResponse{}, errors.New("Validation Failed")
	}

	res, err := server.EidUsed(context.Background(), &proto.EidUsedRequest{Eid: eid})
	if err != nil {
		log.Println("Error Returned from ID used", err.Error())
		return &proto.AuthResponse{}, errors.New("Something went wrong")
	}
	if res.GetUsed() {
		return &proto.AuthResponse{}, errors.New("Given Employee Id already exists")
	}

	newUser := global.User{ID: primitive.NewObjectID(), Eid: eid, Empname: empname, Elevel: elevel, Estream: estream}

	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	_, err = global.DB.Collection("user").InsertOne(ctx, newUser)
	if err != nil {
		log.Println("Error inserting a new user", err.Error())
		return &proto.AuthResponse{}, errors.New("Something went wrong")
	}
	return &proto.AuthResponse{Token: newUser.GetToken()}, nil

}

func (crudServer) EidUsed(_ context.Context, in *proto.EidUsedRequest) (*proto.UsedResponse, error) {
	eid := in.GetEid()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var result global.User
	global.DB.Collection("user").FindOne(ctx, bson.M{"eid": eid}).Decode(&result)
	return &proto.UsedResponse{Used: result != global.NilUser}, nil
}

func (crudServer) AuthUser(_ context.Context, in *proto.AuthUserRequest) (*proto.AuthUserResponse, error) {
	token := in.GetToken()
	user := global.UserFromToken(token)
	return &proto.AuthUserResponse{Eid: user.Eid, Empname: user.Empname, Elevel: user.Elevel, Estream: user.Estream}, nil
}

//Create Operation ends

func (crudServer) Retrieve(_ context.Context, in *proto.RetrieveRequest) (*proto.AuthResponse, error) {
	eid := in.GetEid()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var user global.User
	global.DB.Collection("user").FindOne(ctx, bson.M{"eid": eid}).Decode(&user)
	if user == global.NilUser {
		return &proto.AuthResponse{}, errors.New("No such Employee Id exists")
	}
	return &proto.AuthResponse{Token: user.GetToken()}, nil
}

//Retrieve Operation Ends

func (crudServer) Update(_ context.Context, in *proto.UpdateRequest) (*proto.AuthResponse, error) {
	eid := in.GetEid()
	empname1 := in.GetEmpname1()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var user global.User
	_, err := global.DB.Collection("user").UpdateOne( //Update One document
		ctx,
		bson.M{"eid": eid},
		bson.D{
			{"$set", bson.D{{"empname", empname1}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return &proto.AuthResponse{Token: user.GetToken()}, nil
}

//Update Operation Ends

func (crudServer) Delete(_ context.Context, in *proto.DeleteRequest) (*proto.AuthResponse, error) {
	eid := in.GetEid()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	filter := bson.M{"eid": eid}
	var user global.User
	res, err := global.DB.Collection("user").DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v Documents! \n", res.DeletedCount)
	return &proto.AuthResponse{Token: user.GetToken()}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterCrudServiceServer(server, crudServer{})
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal("Error creating Listener")
	}
	go func() {
		log.Fatal("serving gRPC: ", server.Serve(listener).Error())
	}()

	grpcWebServer := grpcweb.WrapServer(server)
	httpServer := &http.Server{
		Addr: ":9090",
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
