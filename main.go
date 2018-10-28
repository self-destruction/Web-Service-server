package main

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "./proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

type server struct{}

func (s *server) GetBankAnswer(ctx context.Context, in *pb.UserRequest) (*pb.BankReply, error) {
	RunesForName := []rune(in.User.Name)
	SubName := string(RunesForName[0])

	RunesForPatronymic := []rune(in.User.Patronymic)
	SubPatronymic := string(RunesForPatronymic[0])

	rand.Seed(time.Now().UnixNano())
	RandBigInt := rand.Intn(500000)
	RandBigPercent := rand.Intn(80)

	TextMessage := "Уважаемый (ая) " + in.User.Surname + " " + SubName + "." + SubPatronymic + ". "
	var ApproveStatus = true

	if in.User.Age >= 18 {
		TextMessage += "Рады сообщить, что Вам одобрен кредит в размере " + strconv.Itoa(RandBigInt) + " рублей под " +
			strconv.Itoa(RandBigPercent) + "% годовых. Ваш любимый ПромФакБанк."
	} else {
		TextMessage += "К сожалению, вынуждены Вам сообщить об отказе в предоставлении кредита. Ваш любимый ПромФакБанк."
		ApproveStatus = false
	}
	return &pb.BankReply{Message: TextMessage, Approved: ApproveStatus}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCreditServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
