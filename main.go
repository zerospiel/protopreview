package main

import "google.golang.org/grpc"

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("recovered")
		}
	}()
	println("hello world")
	conn, err := grpc.Dial("0.0.0.0:82", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()
}
