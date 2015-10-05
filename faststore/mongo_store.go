// Copyright 2015 Benjamin Campbell <benji@benjica.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package faststore

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// A URLModel abstracts the fields.
type URLModel struct {
	ID   bson.ObjectId `bson:"_id"`
	Slug string        `bson:"slug"`
	URL  string        `bson:"url"`
}

// A URLMongoStore represents a persistent connection to a mongo backend.
type URLMongoStore struct {
	session    *mgo.Session
	collection *mgo.Collection
}

// NewURLMongoStore is a factory function to return a valid URLMongoStore.
func NewURLMongoStore(uri string) (s *URLMongoStore, err error) {
	s = &URLMongoStore{}
	s.session, err = mgo.Dial(uri)
	if err != nil {
		return s, err
	}

	s.collection = s.session.DB("local").C("urls")
	return s, err
}

// AddURL adds the given URLModel to the mongodb
func (s *URLMongoStore) AddURL(model URLModel) {
	model.ID = bson.NewObjectId()

	err := s.collection.Insert(model)
	if err != nil {
		panic(err)
	}
}

// GetURL attempts to find a record with the given slug. If no record is found
// an error is returned.
func (s *URLMongoStore) GetURL(id string) (model URLModel, err error) {

	if err = s.collection.Find(bson.M{"slug": id}).One(&model); err != nil {
		return model, errors.New("Did not find URL with given id")
	}

	return model, err
}

// AllURLs aggregates all model records.
func (s *URLMongoStore) AllURLs() (models []URLModel, err error) {
	if err = s.collection.Find(nil).All(&models); err != nil {
		return models, err
	}

	return models, nil
}
