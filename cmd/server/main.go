package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/unbeatable-abayomi/secondnewPractiseTutorialEdge/internal/transport/http"
)

//App - the struct which contains things like pointers to database connections

type App struct {

}

//RUN sets up our application
func (app *App) Run() error{
	fmt.Println("Setting up Our App")

handler := transportHTTP.NewHandler()
handler.SetupRoutes()
if err := http.ListenAndServe(":8080", handler.Router); err != nil{
	fmt.Println("Failed to setup server")
	return err
}
	return nil
}

func main(){
	fmt.Println("Rest Api Course")

	app := App{};

	if err := app.Run(); err != nil{
		fmt.Println("Error in starting the application");
		fmt.Println(err);
	}
}