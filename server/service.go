package server

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"grpc-etcd/pb"
	"log"
	"net"
	"os"
	"os/signal"
)

var c = make(chan os.Signal)

type student struct {
	pb.FindTeacherServer
}

func (s *student) SayTeacher(c context.Context, stu *pb.Student) (*pb.Teacher, error) {
	if stu.Name == "sa" && stu.Age == 23 {
		return &pb.Teacher{
			Name: "alex",
		}, nil
	} else {
		return nil, errors.New("don't have this student")
	}
}

func CreateService() {
	prefix, addr := "/etcd", "192.168.1.51:8800"
	signal.Notify(c)

	srv, err := NewServiceRegister(prefix, addr, 5)
	if err != nil {
		log.Println(err)
	}
	go srv.ListenKeepChan()

	ser := grpc.NewServer()
	pb.RegisterFindTeacherServer(ser, new(student))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	log.Println("grpc server start~~~")

	go func() {
		for {
			select {
			case <-c:
				err := srv.Close()
				if err != nil {
					log.Println("exit server error: ", err.Error())
				}
				os.Exit(1)
			}
		}
	}()
	_ = ser.Serve(listener)
}
