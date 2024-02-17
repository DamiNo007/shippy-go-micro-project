package main

import (
	"context"
	"gopkg.in/mgo.v2"
	pb "vessel-service/proto/vessel"
)

type handler struct {
	session *mgo.Session
}

func (h *handler) GetRepo() Repository {
	return &VesselRepository{h.session.Clone()}
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := h.GetRepo()

	defer repo.Close()

	vessel, err := repo.FindAvailable(req)

	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := h.GetRepo()

	defer repo.Close()

	if err := repo.Create(req); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true

	return nil
}
