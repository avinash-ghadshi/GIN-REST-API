package main

func main() {
	app := setupRouter()
	app.Run("localhost:9090")
}
