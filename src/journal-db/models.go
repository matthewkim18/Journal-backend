package main

import (
	"github.com/globalsign/mgo/bson"
)

type Journal struct {
	ID      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Subject string        `json:"subject" bson:"subject"`
	Text    string        `json:"text" bson:"text"`
}

// A map to store the books with the ISDN as the key
// This acts as the storage in lieu of an actual database
var journal_list = make(map[string]Journal)
