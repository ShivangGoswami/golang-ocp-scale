package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ShivangGoswami/golang-ocp-scale/async-time-service/appConfig"
	messageSrv "github.com/ShivangGoswami/golang-ocp-scale/async-time-service/messageService"
	timeSrv "github.com/ShivangGoswami/golang-ocp-scale/async-time-service/timeService"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
		//log.Printf("request:%#v", r)
		fmt.Fprintf(rw, "%s --- %s", timeSrv.GetTime().Format("_2-Jan-2006, 3:04:05 PM"), messageSrv.GetMessage())
	}).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))

	srv := &http.Server{
		Addr: appConfig.Config.AppUrl,
		// WriteTimeout: time.Second * 15,
		// ReadTimeout:  time.Second * 15,
		// IdleTimeout:  time.Second * 60,
		Handler: r,
	}

	//server
	go func() {
		log.Println("Server is Starting......")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	//scheduler
	go func() {
		time.Sleep(appConfig.Config.SchedulerDelay)
		for {
			response, err := http.Get("http://" + appConfig.Config.AppUrl + "/time")
			if err != nil {
				fmt.Println("Error:", err)
			}
			parsedresponse, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(parsedresponse))
			time.Sleep(appConfig.Config.SchedulerInterval)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
