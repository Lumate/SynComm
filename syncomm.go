//
// Name: SynComm
// Desc: Zeromq Rebroadcasting server
// Purpose: To allow synchronization of data across the system
//     and allow aggregation of messages for statistics
//
// Author: James Marlowe
//
package main

import (
    "fmt"
    "strings"
    zmq "github.com/alecthomas/gozmq"
)

func main() {
    context, _ := zmq.NewContext()
    receiver, _ := context.NewSocket(zmq.REP)
    publisher, _ := context.NewSocket(zmq.PUB)
    defer context.Close()
    defer receiver.Close()
    defer publisher.Close()

    fmt.Printf("Ready to publish...")
    publisher.Bind("tcp://*:5556")
    
    fmt.Printf("Listening...")
    receiver.Bind("tcp://*:5555")

    for {
        receive, _ := receiver.Recv(0)
        received := string(receive[:])
        receiver.Send([]byte("ok"), 0)
        // messages received with "|" will be treated as <channel>|<message>
        if strings.Contains(received, "|") {
            split_msg := strings.SplitN(received, "|", 2)
            publisher.SendMultipart([][]byte{[]byte(split_msg[0]), []byte(split_msg[1])}, 0)
        } else {
            publisher.Send([]byte(received), 0)
        }
    }
}
