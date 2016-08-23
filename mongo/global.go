package mongo

import (
	"gopkg.in/mgo.v2"
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

// Session get current session
func Session() *mgo.Session {
	return GH.Session()
}

// CloneSession clone session
func CloneSession() *mgo.Session {
	return GH.CloneSession()
}

// DB get database
func DB() *mgo.Database {
	return GH.DB()
}

// DBWithName get database with name
func DBWithName(dbName string) *mgo.Database {
	return GH.DBWithName(dbName)
}

// CHandle use the default DB processing with a set of callback functions,
// the session will be closed down after processing is completed
func CHandle(cName string, handle func(c *mgo.Collection)) {
	GH.CHandle(cName, handle)
}

// CHandleWithDB use the new DB processing with a set of callback functions,
// the session will be closed down after processing is completed
func CHandleWithDB(dbName, cName string, handle func(c *mgo.Collection)) {
	GH.CHandleWithDB(dbName, cName, handle)
}

// C use the default db get collection
func C(cName string) *mgo.Collection {
	return GH.C(cName)
}

// CWithDB use the new db get collection
func CWithDB(dbName, cName string) *mgo.Collection {
	return GH.CWithDB(dbName, cName)
}

// IncrID get increase id
func IncrID(cName string, storeCName ...string) (id int64, err error) {
	id, err = GH.IncrID(cName, storeCName...)
	return
}
