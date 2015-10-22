package config

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/sessions"
)

type SessionData struct{
	store *sessions.CookieStore;
}

var Session *SessionData;

func init(){
	Session = &SessionData{};
	Session.store = sessions.NewCookieStore([]byte("something-very-secret"));
}

func (this *SessionData) Get(c *gin.Context) (map[interface{}]interface{},error){
	session, err := this.store.Get(c.Request, "gosession");
	return session.Values,err;
}

func (this *SessionData) Set(c *gin.Context,key interface{},value interface{}) (error){
	session, err := this.store.Get(c.Request, "gosession");
	if err != nil {
		return err;
	}
	session.Values[key] = value;
	return session.Save(c.Request,c.Writer);
}

func (this *SessionData) Destroy(c *gin.Context) (error){
	session, err := this.store.Get(c.Request, "gosession");
	if err != nil {
		return err;
	}
	for key,_ := range session.Values{
		delete(session.Values,key);
	}
	return session.Save(c.Request,c.Writer);
}