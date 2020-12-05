package main

import (
  "github.com/myapp/route"
)

func main() {
    // Routes
    router := route.Init()
    router.Logger.Fatal(router.Start(":1323"))
}