package main

import (
	"context"
	"flag"
	"io"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"

	pb "github.com/ngenator/tagbot/service/dice/dice"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func roll(client pb.RollerClient, dice *pb.Dice) []int32 {
	log.Printf("Requesting dice roll with rolls=%d, sides=%d", dice.Rolls, dice.Sides)
	stream, err := client.Roll(context.Background(), dice)
	if err != nil {
		log.Fatalf("%v.Roll(_) = _, %v", client, err)
	}

	results := []int32{}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			// read done.
			return results
		}
		if err != nil {
			log.Fatalf("Failed to receive results : %v", err)
		}
		log.Printf("Roll %d: %d", in.Roll, in.Result)
		results = append(results, in.Result)
	}

	return results
}

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRollerClient(conn)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 10; i++ {
		results := roll(client, &pb.Dice{
			Rolls: r.Int31n(5) + 1,
			Sides: r.Int31n(20) + 1,
		})
		log.Println(results)
	}
}
