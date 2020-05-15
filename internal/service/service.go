package service

import (
	"context"
	"fmt"

	pb "order/api"
	"order/internal/dao"
	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply = new(empty.Empty)
	fmt.Printf("hello %s", req.Name)
	return
}

// SayHelloURL bm demo func.
func (s *Service) SayHelloURL(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	reply = &pb.HelloResp{
		Content: "hello " + req.Name,
	}
	fmt.Printf("hello url %s", req.Name)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}

// Create
func (s *Service) Create(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Create" + req.Name,
	}
	fmt.Printf("Create %s", req.Name)
	return
}

// Delete
func (s *Service) Delete(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Delete" + req.Name,
	}
	fmt.Printf("Delete %s", req.Name)
	return
}

// Get
func (s *Service) Get(ctx context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	reply = &pb.Resp{
		Content: "Get" + req.Name,
	}

	//s.dao.Article(c,req.XXX_sizecache)

	fmt.Printf("Get %s", req.Name)
	return
}

func (s *Service) TestQ(c context.Context, req *pb.Req) (reply *pb.Resp, err error) {
	aa, err := s.dao.Article(c, 1)
	if err != nil {
		return
	}
	reply = &pb.Resp{
		Content: "123",
	}
	fmt.Println(aa,err,"*****")
	return
}
