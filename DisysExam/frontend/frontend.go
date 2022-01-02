package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	pb "github.com/CasperAntonPoulsen/DisysExam/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Replica struct {
	user       *pb.User
	port       string
	connection pb.ExamClient
}

type Server struct {
	pb.UnimplementedExamServer
	leader   Replica
	Replicas []Replica
}

func (s *Server) FEIncrement(c *gin.Context) {

	log.Print("Getting current bid")
	rqst := &pb.Request{User: &pb.User{Userid: 0}}

	stream, err := s.leader.connection.RequestToken(context.Background(), rqst)
	if err != nil {
		// Throw and internal server error message and send the exception in the ack
		c.JSON(500, gin.H{
			"ack": fmt.Sprintln(err),
		})
		return
	}

	//recieve grant token

	for {
		_, err := stream.Recv()
		if err != nil {
			log.Printf("Error recieving message: %v", err)
			break
		} else {
			log.Print("Grant token recieved, accessing critical section")
			// access critical section

			result, err := s.leader.connection.Increment(context.Background(), &pb.Empty{})
			if err != nil {
				// Throw and internal server error message and send the exception in the ack
				c.JSON(500, gin.H{
					"ack": fmt.Sprintln(err),
				})
				return
			}

			log.Printf("Result: %v", result)

			c.String(http.StatusOK, strconv.Itoa(int(result.Amount)))

			log.Print("Finished, sending release token")
			// then release
			s.leader.connection.ReleaseToken(context.Background(), &pb.Release{User: rqst.User})
			break
		}
	}

}

func (s *Server) Coordinator(ctx context.Context, state *pb.State) (*pb.Empty, error) {
	for _, rm := range s.Replicas {
		if rm.user.Userid == state.User.Userid {
			s.leader = rm
		}
	}
	log.Printf("Recivied new leader: %v", s.leader.user.Userid)
	return &pb.Empty{}, nil
}

//helper function to convert string to int32
func GetIntEnv(envvar string) int32 {
	envvarstring := os.Getenv(envvar)
	envvarint, err := strconv.Atoi(envvarstring)
	if err != nil {
		log.Fatalf("not a valid int: %v", err)
	}

	return int32(envvarint)
}

func main() {

	grpcServer := grpc.NewServer()

	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Error, couldn't create the server %v", err)
	}

	NumReplicas := GetIntEnv("NREPLICATIONMANAGERS")
	var Replicas []Replica

	for i := 0; i < int(NumReplicas); i++ {

		portInt := 8080 + i
		port := ":" + strconv.Itoa(portInt)

		conn, err := grpc.Dial("incrementserver"+strconv.Itoa(i+1)+port, grpc.WithInsecure())
		if err != nil {
			log.Printf("could not connect to rm %v: %v", i+1, err)
		}
		log.Printf("Connected to Replica: %v", i+1)
		rmClient := pb.NewExamClient(conn)

		Replicas = append(Replicas, Replica{&pb.User{Userid: int32(i + 1)}, port, rmClient})
	}

	server := Server{
		leader:   Replicas[0],
		Replicas: Replicas,
	}

	pb.RegisterExamServer(grpcServer, &server)
	r := gin.Default()

	r.GET("/Increment", server.FEIncrement)

	go func() { grpcServer.Serve(listener) }()
	log.Print("API listening at :8000")
	r.Run(":8000")
}
