package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// InitHandler 初始化Handler
// url MongoDB连接路径
func InitHandler(url string) (*Handler, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Handler{
		session: session,
	}, nil
}

// Handler 提供MongoDB的数据库操作
type Handler struct {
	session *mgo.Session
}

// Session 获取当前会话
func (h *Handler) Session() *mgo.Session {
	return h.session
}

// CloneSession 克隆一个会话
func (h *Handler) CloneSession() *mgo.Session {
	return h.session.Clone()
}

// DBHandle 使用新的DB处理带有回调函数的集合,处理完成之后将关闭该会话(处理时将使用克隆的会话)
func (h *Handler) DBHandle(dbName, cName string, handle func(c *mgo.Collection)) {
	session := h.CloneSession()
	defer session.Close()
	handle(session.DB(dbName).C(cName))
}

// CHandle 使用默认DB处理带有回调函数的集合,处理完成之后将关闭该会话(处理时将使用克隆的会话)
func (h *Handler) CHandle(cName string, handle func(c *mgo.Collection)) {
	session := h.CloneSession()
	defer session.Close()
	handle(session.DB("").C(cName))
}

// DB 使用新的DB处理集合(使用当前会话)
func (h *Handler) DB(dbName, cName string) *mgo.Collection {
	return h.session.DB(dbName).C(cName)
}

// C 使用默认DB处理集合(使用当前会话)
func (h *Handler) C(cName string) *mgo.Collection {
	return h.session.DB("").C(cName)
}

// IncrID 返回一个自增ID
// cName 需要生成自增ID的集合名称
// storeCName 存储自增ID的集合名(默认为counters)
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
