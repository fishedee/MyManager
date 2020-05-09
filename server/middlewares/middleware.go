package middlewares

import (
	. "github.com/fishedee/app/metric"
	. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	"net/http"
	"time"
)

var globalMetric Metric

func NothingDumpMiddleware(oldHandler http.HandlerFunc) http.HandlerFunc {
	log := GetAppBasic().Log
	return func(response http.ResponseWriter, request *http.Request) {
		url := request.URL.Path
		log.Debug("Request in %v", url)
		oldHandler(response, request)
		log.Debug("Request out %v", url)
	}
}

func AllQuestMiddleware(oldHandler http.HandlerFunc) http.HandlerFunc {
	serverRequest := globalMetric.GetTimer("server.request")
	return func(response http.ResponseWriter, request *http.Request) {
		begin := time.Now()
		oldHandler(response, request)
		duration := time.Now().Sub(begin)
		GetAppBasic().Log.Debug("serverRequest update!")
		serverRequest.Update(duration)
	}
}

func PathQuestMiddleware(oldHandler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		url := request.URL.Path
		urlSeg := Explode(url, "/")
		firstSeg := urlSeg[0]

		pathRequest := globalMetric.GetTimer("path.request?ao=" + firstSeg)
		begin := time.Now()
		oldHandler(response, request)
		duration := time.Now().Sub(begin)
		pathRequest.Update(duration)
	}
}

func initMetric() {
	config := GetAppBasic().Config
	connectUrl := config.GetString("metricconnecturl")
	database := config.GetString("metricdatabase")
	user := config.GetString("metricuser")
	password := config.GetString("metricpassword")

	var err error
	globalMetric, err = NewMetric(MetricConfig{
		ConnectUrl: connectUrl,
		Database:   database,
		User:       user,
		Password:   password,
	})
	if err != nil {
		panic(err)
	}
	go globalMetric.Run()
}

func init() {
	initMetric()
	InitRouteMiddleware(NothingDumpMiddleware)
	InitRouteMiddleware(AllQuestMiddleware)
	InitRouteMiddleware(PathQuestMiddleware)
}
