package main

import (
	"encoding/json"
	"errors"
	"go.unistack.org/micro/v3/broker"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
	pb "user-service/proto/user"
)

type handler struct {
	repo         Repository
	tokenService Authable
	PubSub       broker.Broker
}

const onUserCreatedTopic = "user.created"

func (srv *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := srv.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token

	return nil
}

func (srv *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	req.Password = string(hashedPass)

	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req

	if err := srv.publishEvent(ctx, req); err != nil {
		return err
	}

	return nil
}

func (srv *handler) publishEvent(ctx context.Context, user *pb.User) error {
	body, err := json.Marshal(user)

	if err != nil {
		return nil
	}

	msg := &broker.Message{
		Header: map[string]string{
			"id":          user.Id,
			"Micro-Topic": onUserCreatedTopic,
		},
		Body: body,
	}

	if err := srv.PubSub.Publish(ctx, onUserCreatedTopic, msg); err != nil {
		log.Printf("[pub] failed %v", err.Error())
	}

	return nil
}

func (srv *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := srv.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println(claims)

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}
