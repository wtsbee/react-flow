package controller

import (
	"log"
	"net/http"
	"react-flow-api/model"
	"react-flow-api/usecase"
	"strconv"

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
	nodes, edges, err := nc.nu.GetAllNodes()
	if err != nil {
		log.Println("controller GetAllNodes ノード取得エラー: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	log.Println("controller GetAllNodes : ノード取得成功")

	// nodeデータの整形
	nodeDetails := make([]model.NodeDetail, len(nodes))
	for i, node := range nodes {
		nodeDetails[i] = model.NodeDetail{
			ID:       strconv.Itoa(int(node.ID)),
			Type:     node.Type,
			Position: model.PositionDetails{X: node.X, Y: node.Y},
			Data:     model.DataDetails{Label: node.Label},
		}
	}

	// edgeデータの整形
	edgeDetails := make([]model.EdgeDetail, len(edges))
	for i, edge := range edges {
		edgeDetails[i] = model.EdgeDetail{
			ID:     strconv.Itoa(int(edge.ID)),
			Source: strconv.Itoa(int(edge.SourceID)),
			Target: strconv.Itoa(int(edge.TargetID)),
		}
	}

	// レスポンスデータの生成
	resData := model.NodeAndEdgeResponse{Node: nodeDetails, Edge: edgeDetails}

	return c.JSON(http.StatusOK, resData)
}
