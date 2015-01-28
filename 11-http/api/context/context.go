package context

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/go-uuid/uuid"
	"github.com/ArdanStudios/gotraining/11-http/api/mongodb"
	"github.com/sqs/mux"
	"gopkg.in/mgo.v2"
)

// Context contains data in context with all requests.
type Context struct {
	Session   *mgo.Session
	Writer    http.ResponseWriter
	Request   *http.Request
	SessionID string
}

// AddRoute allows routes to be injected into the middleware with the context.
func AddRoute(router *mux.Router, path string, userHandler func(c *Context)) {
	f := func(w http.ResponseWriter, r *http.Request) {
		uid := uuid.New()
		log.Printf("%s : context : handler : Started : Path[%s] URL[%s]\n", uid, path, r.URL.RequestURI())

		c := Context{
			Writer:    w,
			Request:   r,
			Session:   mongodb.GetSession(),
			SessionID: uid,
		}

		defer func() {
			if r := recover(); r != nil {
				log.Println(uid, ": context : handler : PANIC :", r)
			}

			c.Session.Close()
			log.Println(uid, ": context : handler : Completed")
		}()

		if err := c.authenticate(); err != nil {
			log.Println(uid, ": context : handler : ERROR :", err)
			return
		}

		userHandler(&c)
	}

	router.HandleFunc(path, f)
	log.Printf("main : context : AddRoute : Added : Path[%s]\n", path)
}

// authenticate handles the authentication of each request.
func (c *Context) authenticate() error {
	log.Println(c.SessionID, ": context : authenticate : Started")

	// ServeError(w, errors.New("Auth Error"), http.StatusUnauthorized)

	log.Println(c.SessionID, ": context : authenticate : Completed")
	return nil
}

// ServeError handles application errors
func (c *Context) ServeError(err error, statusCode int) {
	log.Printf("%s : context : ServeError : Started : Error[%s]\n", c.SessionID, err)

	e := struct {
		Err string
	}{
		Err: err.Error(),
	}

	data, err := json.MarshalIndent(&e, "", "    ")
	if err != nil {
		http.Error(c.Writer, "", http.StatusInternalServerError)
		return
	}

	http.Error(c.Writer, string(data), statusCode)
	log.Println(c.SessionID, ": context : ServeError : Completed")
}

// ServeJSON handles serving values as JSON.
func (c *Context) ServeJSON(v interface{}) {
	log.Printf("%s : context : ServeJSON : Started : Error[%+v]\n", c.SessionID, v)

	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		http.Error(c.Writer, "", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(c.Writer, string(data))
	log.Println(c.SessionID, ": context : ServeJSON : Completed")
}
