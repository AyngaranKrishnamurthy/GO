package main

import (
	"context"
	"go_blog/crud_func/global"
	"go_blog/crud_func/proto"
	"testing"
)

func Test_crudServer_CreateEmp(t *testing.T) {
	global.ConnectToTestDB()
	global.DB.Collection("user").InsertOne(context.Background(), global.User{EID: "13246170", Empname: "Ayngaran", Level: "12", Stream: "Java"})
	server := crudServer{}
	_, err := server.CreateEmp(context.Background(), &proto.CreateEmpRequest{Id: "13246170", Name: "Ayngaran", Level: "12", Stream: "Java"})
	if err != nil {
		t.Error("Error occured at C")
	}

}

func Test_crudServer_ReadEmp(t *testing.T) {
	global.ConnectToTestDB()
	//global.DB.Collection("user").InsertOne(context.Background(), global.User{EID: "13246170", Empname: "Ayngaran", Level: "12", Stream: "Java"})
	server := crudServer{}
	_, err := server.ReadEmp(context.Background(), &proto.ReadEmpRequest{Id: "13246170"})
	if err != nil {
		t.Error("Error Occured at R")
	}
}

func Test_crudServer_UpdateEmp(t *testing.T) {
	global.ConnectToTestDB()
	//global.DB.Collection("user").InsertOne(context.Background(), global.User{EID: "13246170", Empname: "Ayngaran", Level: "12", Stream: "Java"})
	server := crudServer{}
	_, err := server.UpdateEmp(context.Background(), &proto.UpdateEmpRequest{Id: "13246170", Nname: "Karan"})
	if err != nil {
		t.Error("Error Occured at U")
	}
}

func Test_crudServer_DeleteEmp(t *testing.T) {
	global.ConnectToTestDB()
	//global.DB.Collection("user").InsertOne(context.Background(), global.User{EID: "13246170", Empname: "Ayngaran", Level: "12", Stream: "Java"})
	server := crudServer{}
	_, err := server.UpdateEmp(context.Background(), &proto.UpdateEmpRequest{Id: "13246170"})
	if err != nil {
		t.Error("Error Occured at D")
	}
}
