package controller

import (
	"log"
	"net/http"
	"react-flow-api/usecase"

	"github.com/labstack/echo/v4"
)

type INodeController interface {
	GetAllNodes(c echo.Context) error
}

type nodeController struct {
	nu usecase.INodeUsecase
}

// コンストラクタ
func NewNodeController(nu usecase.INodeUsecase) INodeController {
	return &nodeController{nu}
}

func (nc *nodeController) GetAllNodes(c echo.Context) error {
	nodes, err := nc.nu.GetAllNodes()
	if err != nil {
		log.Println("controller GetNode ノード取得エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller GetNode : ノード取得成功")
	return c.JSON(http.StatusOK, nodes)
}
