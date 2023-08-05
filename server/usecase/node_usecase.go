package usecase

import (
	"react-flow-api/model"
	"react-flow-api/repository"
)

type INodeUsecase interface {
	GetAllNodes() ([]model.Node, []model.Edge, error)
}

type nodeUsecase struct {
	nr repository.INodeRepository
	er repository.IEdgeRepository
}

// コンストラクタ
func NewNodeUsecase(nr repository.INodeRepository, er repository.IEdgeRepository) INodeUsecase {
	return &nodeUsecase{nr, er}
}

func (nu *nodeUsecase) GetAllNodes() ([]model.Node, []model.Edge, error) {
	var nodes []model.Node
	if err := nu.nr.GetAllNodes(&nodes); err != nil {
		return nil, nil, err
	}
	var edges []model.Edge
	if err := nu.er.GetAllEdges(&edges); err != nil {
		return nil, nil, err
	}
	return nodes, edges, nil
}
