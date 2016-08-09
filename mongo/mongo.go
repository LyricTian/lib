package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// define global handler
var (
	GH *Handler
)

// LoadHandler load global handler
func LoadHandler(url, dbName string) {
	h, err := InitHandlerWithDB(url, dbName)
	if err != nil {
		panic(err)
	}
	GH = h
}

// InitHandler initialize the Handler
func InitHandler(url string) (*Handler, error) {
	return InitHandlerWithDB(url, "")
}

// InitHandlerWithDB initialize the Handler with db name
func InitHandlerWithDB(url, dbName string) (*Handler, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Handler{
		session: session,
		dbName:  dbName,
	}, nil
}

// Handler Provide the mongo's operation
type Handler struct {
	session *mgo.Session
	dbName  string
}

// Session get current session
func (h *Handler) Session() *mgo.Session {
	return h.session
}

// CloneSession clone session
func (h *Handler) CloneSession() *mgo.Session {
	return h.session.Clone()
}

// DB get database
func (h *Handler) DB() *mgo.Database {
	return h.session.DB(h.dbName)
}

// DBWithName get database with name
func (h *Handler) DBWithName(dbName string) *mgo.Database {
	return h.session.DB(dbName)
}

// CHandle use the default DB processing with a set of callback functions,
// the session will be closed down after processing is completed
func (h *Handler) CHandle(cName string, handle func(c *mgo.Collection)) {
	session := h.CloneSession()
	defer session.Close()
	handle(session.DB(h.dbName).C(cName))
}

// CHandleWithDB use the new DB processing with a set of callback functions,
// the session will be closed down after processing is completed
func (h *Handler) CHandleWithDB(dbName, cName string, handle func(c *mgo.Collection)) {
	session := h.CloneSession()
	defer session.Close()
	handle(session.DB(dbName).C(cName))
}

// C use the default db get collection
func (h *Handler) C(cName string) *mgo.Collection {
	return h.session.DB(h.dbName).C(cName)
}

// CWithDB use the new db get collection
func (h *Handler) CWithDB(dbName, cName string) *mgo.Collection {
	return h.session.DB(dbName).C(cName)
}

// IncrID get increase id
func (h *Handler) IncrID(cName string, storeCName ...string) (id int64, err error) {
	sCName := "counters"
	if len(storeCName) > 0 {
		sCName = storeCName[0]
	}
	h.CHandle(sCName, func(c *mgo.Collection) {
		var result struct {
			Seq int64 `bson:"seq"`
		}
		_, err = c.Find(bson.M{"_id": cName}).Apply(mgo.Change{
			Update:    bson.M{"$inc": bson.M{"seq": 1}},
			ReturnNew: true,
			Upsert:    true,
		}, &result)
		if err == nil {
			id = result.Seq
		}
	})
	return
}
