package router

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/sessions"
    "net/http"
    "time"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"));

func SetLoginRouter(router *gin.RouterGroup){
	router.GET("/islogin", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "gosession");

		if userId , ok := session.Values["userId"] ; ok {
			c.String(http.StatusOK, "haslogin " + userId.(string));
		}else{
			c.String(http.StatusOK, "has not login");
		}
	});
	router.GET("/checkin", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "gosession");
		session.Values["userId"] = time.Now().String();
		session.Save(c.Request,c.Writer);

		c.String(http.StatusOK, "checkin");
	});
	router.GET("/checkout", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "gosession");
		delete( session.Values , "userId");
		session.Save(c.Request,c.Writer);

		c.String(http.StatusOK, "checkout");
	});
}