// Copyright (C) 2022 Andreas P. <me@apap04.com>

// This file is part of the shortlink project.

// The shortlink project can not be copied and/or distributed without the express
// permission of Andreas P. <me@apap04.com>.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/uptrace/bunrouter"
)

type Shortlink struct {
	// ID   string `json:id`
	Link string `json:link`
	// Creator string `json:creator`
}

func (m Shortlink) UnmarshalBinary(data []byte) error {
	// convert data to yours, let's assume its json data
	return json.Unmarshal(data, m)
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found. falling back")
	}

	router := bunrouter.New()

	fs := http.FileServer(http.Dir("./files"))

	router.GET("/", bunrouter.HTTPHandler(fs))

	router.GET("/:id", bunrouter.HTTPHandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		params := bunrouter.ParamsFromContext(req.Context())

		rdb := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASS"),
			DB:       0,
		})

		val, err := rdb.Get("shortlink-" + params.ByName("id")).Result()
		if err != nil {
			w.Write([]byte("Not found"))
			return
		}

		defer rdb.Close()
		http.Redirect(w, req, val, http.StatusFound)
	}))

	router.POST("/api/create", bunrouter.HTTPHandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		var jsonreq Shortlink
		err := json.NewDecoder(req.Body).Decode(&jsonreq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rdb := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASS"),
			DB:       0,
		})
		id, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 7)
		if err != nil {
			fmt.Println(err)
		}
		err = rdb.Set("shortlink-"+id, jsonreq.Link, 0).Err()
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write([]byte("yeat.dev/" + id))
	}))

	log.Println("Server started on port " + os.Getenv("PORT"))
	log.Println(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func main() {
	Run()
}
