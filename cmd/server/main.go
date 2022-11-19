package main


import "fmt"
  
//App - the struct which contains things like pointers to database connections

type App struct {

}

//RUN sets up our application
func (app *App) Run() error{
	fmt.Println("Setting up Our App")

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