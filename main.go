// Copyright (C) 2022 Andreas P. <me@apap04.com>

// This file is part of the shortlink project.

// The shortlink project can not be copied and/or distributed without the express
// permission of Andreas P. <me@apap04.com>.

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
	"yeat.dev/shortlink/handlers"
)

var ctx = context.Background()

func debugHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return bunrouter.JSON(w, bunrouter.H{
		"route":  req.Route(),
		"params": req.Params().ByName("id"),
	})
}

func main() {
	router := bunrouter.New()

	router.WithGroup("/s", func(g *bunrouter.Group) {
		g.GET("/:id", handlers.GotoLink)
	})

	log.Println("Server started on port 9999")
	log.Println(http.ListenAndServe(":9999", router))
}
