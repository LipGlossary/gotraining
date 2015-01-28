// Package mongodb provides driver support.
package mongodb

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

// MongoDB connection information.
const (
	mongoDBHosts = "ds039441.mongolab.com:39441"
	authDatabase = "gotraining"
	authUserName = "got"
	authPassword = "got2015"
	testDatabase = "gotraining"
)

// DBCall defines a type of function that can be used
// to excecute code against MongoDB.
type DBCall func(*mgo.Collection) error

// session maintains the master session
var session *mgo.Session

// init sets up the MongoDB environment.
func init() {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := mgo.DialInfo{
		Addrs:    []string{mongoDBHosts},
		Timeout:  60 * time.Second,
		Database: authDatabase,
		Username: authUserName,
		Password: authPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	var err error
	if session, err = mgo.DialWithInfo(&mongoDBDialInfo); err != nil {
		log.Fatalln("MongoDB Dial", err)
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)
}

// Log provides a string version of the value
func Log(value interface{}) string {
	json, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(json)
}

// GetSession returns a copy of the master session for use.
func GetSession() *mgo.Session {
	return session.Copy()
}

// Execute the MongoDB literal function.
func Execute(session *mgo.Session, collectionName string, dbCall DBCall) error {
	log.Printf("Execute : Started : Collection[%s]\n", collectionName)

	// Capture the specified collection.
	collection := session.DB(testDatabase).C(collectionName)
	if collection == nil {
		err := fmt.Errorf("Collection %s does not exist", collectionName)
		log.Println("Execute : ERROR :", err)
		return err
	}

	// Execute the MongoDB call.
	err := dbCall(collection)
	if err != nil {
		log.Println("Execute : ERROR :", err)
		return err
	}

	log.Println("Execute Completed")
	return nil
}
