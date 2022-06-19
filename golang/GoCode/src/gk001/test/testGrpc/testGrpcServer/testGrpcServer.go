package main

import (
	"context"
	"net/rpc"
)

type SearchService struct {}

func (s *SearchService)Search(ctx context.Context,r *pb.SearchRequest)


func main(){
	rpc.RegisterName()


}

