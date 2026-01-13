package main

import "gopring/app"

func main() {
	e := app.Bootstrap()
	e.Logger.Fatal(e.Start(":8080"))
}
