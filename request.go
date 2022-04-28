package main

import (
	"grpc-etcd/client"
	"log"
)

func main() {
	teacher, err := client.FindTeacher()
	if err != nil {
		log.Println(err)
	}
	log.Println(teacher.String())
}
