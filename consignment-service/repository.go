package main

import (
	pb "consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	mongoSession *mgo.Session
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.mongoSession.DB(dbName).C(consignmentCollection)
}

func (repo *ConsignmentRepository) Close() {
	repo.mongoSession.Close()
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}
