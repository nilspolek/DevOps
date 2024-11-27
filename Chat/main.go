package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nilspolek/DevOps/Chat/direct_message_service/dm_impl"
	dmlog "github.com/nilspolek/DevOps/Chat/direct_message_service/dm_log"
	dmprometheus "github.com/nilspolek/DevOps/Chat/direct_message_service/dm_prometheus"
	_ "github.com/nilspolek/DevOps/Chat/docs"
	gmimpl "github.com/nilspolek/DevOps/Chat/group_message_service/gm_impl"
	gmlog "github.com/nilspolek/DevOps/Chat/group_message_service/gm_log"
	gmprometheus "github.com/nilspolek/DevOps/Chat/group_message_service/gm_prometheus"
	gimpl "github.com/nilspolek/DevOps/Chat/group_service/g_impl"
	glog "github.com/nilspolek/DevOps/Chat/group_service/g_log"
	gprometheus "github.com/nilspolek/DevOps/Chat/group_service/g_prometheus"
	jwtimpl "github.com/nilspolek/DevOps/Chat/jwt_service/jwt_impl"
	"github.com/nilspolek/DevOps/Chat/repo/mongodb"
	"github.com/nilspolek/DevOps/Chat/transport/rest"
	"github.com/nilspolek/goLog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	DEFAULT_PORT      = 8080
	ENABLE_LOG        = true
	ENABLE_PROMETHEUS = false
	ENABLE_SWAGGER    = true
)

// @title Chat API
//
//go:generate swag init
func main() {
	// Assinging the address to the environment variable CHAT_ADDRESS
	var address string
	if address = os.Getenv("CHAT_ADDRESS"); address == "" {
		address = fmt.Sprintf(":%d", DEFAULT_PORT)
	}

	// Setup the repositories and services
	repositories, err := mongodb.New()
	if err != nil {
		log.Panic(err)
	}
	dms := dm_impl.New(repositories)
	gms := gmimpl.New(repositories)
	gs := gimpl.New(repositories)
	jwt := jwtimpl.New()
	mux := mux.NewRouter()

	// Enable logging if enabled
	if ENABLE_LOG {
		dms = dmlog.New(&dms)
		gms = gmlog.New(&gms)
		gs = glog.New(&gs)
	}

	// Enable prometheus if enabled
	if ENABLE_PROMETHEUS {
		dms, err = dmprometheus.New(&dms, "direct_message_service")
		if err != nil {
			panic(err)
		}
		gms, err = gmprometheus.New(&gms, "group_message_service")
		if err != nil {
			panic(err)
		}
		gs, err = gprometheus.New(&gs, "group_service")
		if err != nil {
			panic(err)
		}
		mux.Handle("/metrics", promhttp.Handler())
	}

	// Run the server
	router := rest.New(mux, &dms, &gms, &gs, &jwt)
	if ENABLE_LOG {
		goLog.Info("Server is running on address %s", address)
	}

	// Run swagger if enabled
	if ENABLE_SWAGGER {
		// generate swagger documentation
		mux.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	}

	goLog.Error("%v", http.ListenAndServe(address, router.Router))
}
