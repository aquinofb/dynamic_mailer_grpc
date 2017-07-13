package services

import(
  "log"

  pb "github.com/aquinofb/dynamicmailer/lib"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

const (
  address     = "172.16.1.146:50051"
  defaultName = "world"
)

func SendEmail(payload pb.Payload) int32 {
  conn := grpcDial()
  defer conn.Close()

  client := pb.NewMessagesClient(conn)
  r, err := client.SendIt(context.Background(), &payload)

  if err != nil {
   log.Fatalf("could not send it: %v", err)
  }

  return r.Status
}

func grpcDial() (conn *grpc.ClientConn) {
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
   log.Fatalf("did not connect: %v", err)
  }

  return conn
}
