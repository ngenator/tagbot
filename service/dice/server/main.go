package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/ngenator/tagbot/service/dice/dice"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type diceServer struct{}

func (d *diceServer) Roll(dice *pb.Dice, stream pb.Roller_RollServer) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	log.Printf("Rolling %dd%d", dice.Rolls, dice.Sides)

	var i int32
	for i = 1; i <= dice.Rolls; i++ {
		result := &pb.Result{
			Result: r.Int31n(dice.Sides) + 1,
			Roll:   int32(i),
		}
		if err := stream.Send(result); err != nil {
			return err
		}
	}

	return nil
}

func newServer() *diceServer {
	d := new(diceServer)
	return d
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRollerServer(grpcServer, newServer())
	grpcServer.Serve(listener)
}
