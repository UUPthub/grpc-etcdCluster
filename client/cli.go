package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-etcd/pb"
	"log"
)

func FindTeacher() (*pb.Teacher, error) {
	key := "/etcd"
	discover, err := NewServiceDiscover()
	if err != nil {
		log.Println(err)
	}
	discover.WatchServer(key)
	serverList := discover.GetServerList()
	if len(serverList) != 0 {
		client, err := grpc.Dial(serverList[0], grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		findTeacherClient := pb.NewFindTeacherClient(client)
		teacher, err := findTeacherClient.SayTeacher(context.Background(), &pb.Student{
			Name: "sa",
			Age:  23,
		})
		if err != nil {
			return nil, err
		}
		return teacher, nil
	} else {
		log.Printf("have not server: %s\n", key)
		return nil, err
	}

}
