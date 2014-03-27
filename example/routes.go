package example

import (
    "net/http"
    "github.com/codegangsta/martini"
)

func init() {
    m := martini.Classic()

    // Requests (new)
    m.Get("/request/query", requestQueryHandler)
    m.Get("/request/list", requestListHandler)

    // Handle this all
    http.Handle("/", m)
}

