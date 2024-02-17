package main

import (
	pb "consignment-service/proto/consignment"
	"context"
	vesselPb "github.com/DamiNo007/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"log"
)

type handler struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselServiceClient
}

func (s *handler) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	err = repo.Create(req)

	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()

	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
