// Copyright (C) 2022 Andreas P. <me@apap04.com>

// This file is part of the shortlink project.

// The shortlink project can not be copied and/or distributed without the express
// permission of Andreas P. <me@apap04.com>.

package handlers

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/uptrace/bunrouter"
)

func GotoLink(w http.ResponseWriter, r bunrouter.Request) error {
	rdb := redis.NewClient(&redis.Options{
		// Replace these soon.
		Addr:     "xx",
		Password: "",
		DB:       0,
	})

	val, err := rdb.Get(r.Params().ByName("id")).Result()
	if err != nil {
		return nil
	}

	defer rdb.Close()
	return http.Redirect(w, r.Request, "https://apap04.com", http.StatusFound)
}
