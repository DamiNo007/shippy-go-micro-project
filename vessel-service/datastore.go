package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

func CreateSession(host string) (*mgo.Session, error) {
	log.Printf("Trying to connect to DB host: %s", host)
	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
