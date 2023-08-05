package main

import (
	"react-flow-api/controller"
	"react-flow-api/database"
	"react-flow-api/repository"
	"react-flow-api/router"
	"react-flow-api/usecase"
)

func main() {
	db := database.NewDB()
	database.Migrate(db)
	nodeRepository := repository.NewNodeRepository(db)
	edgeRepository := repository.NewEdgeRepository(db)
	nodeUsecase := usecase.NewNodeUsecase(nodeRepository, edgeRepository)
	nodeController := controller.NewNodeController(nodeUsecase)
	e := router.NewRouter(nodeController)
	e.Logger.Fatal(e.Start(":8080"))
}
