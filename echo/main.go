package main

func main() {
	router := newRouter()
	router.Logger.Fatal(router.Start("localhost:8080"))
}
