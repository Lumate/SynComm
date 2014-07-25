//
// SynComm tests
//
// Author: James Marlowe
//
package main

import (
    "fmt"
    "testing"
    zmq "github.com/alecthomas/gozmq"
)

func TestImports(t *testing.T) {
    fmt.Print("imports probably worked")
    context, _ := zmq.NewContext()
    defer context.Close()
    return
}

//Havent gotten this test to work yet
func TestSockets(t *testing.T) {
    context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.REQ)
    receiver, _ := context.NewSocket(zmq.REP)
    publisher, _ := context.NewSocket(zmq.PUB)
    subscriber, _ := context.NewSocket(zmq.SUB)
    defer context.Close()
    defer socket.Close()
    defer receiver.Close()
    defer publisher.Close()
    defer subscriber.Close()
    socket.Connect("tcp://localhost:5555")
    receiver.Bind("tcp://*:5555")
    publisher.Bind("tcp://*:5556")
    subscriber.Bind("tcp://localhost:5556")
    subscriber.SetSubscribe("")
    msg := fmt.Sprintf("Hello")
    socket.Send([]byte(msg), 0)
    new_msg, _ := receiver.Recv(0)
    fmt.Print(new_msg)
    reply := fmt.Sprintf("World")
    fmt.Print(reply)
    receiver.Send([]byte(reply), 0)
    publisher.Send([]byte(new_msg), 0)
    counter, _ := subscriber.Recv(0)
    fmt.Print(counter)
    return
}
