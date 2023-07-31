package usecase

import (
	"react-flow-api/model"
	"react-flow-api/repository"
)

type INodeUsecase interface {
	GetAllNodes() ([]model.Node, error)
}

type nodeUsecase struct {
	nr repository.INodeRepository
}

// コンストラクタ
func NewNodeUsecase(nr repository.INodeRepository) INodeUsecase {
	return &nodeUsecase{nr}
}

func (nu *nodeUsecase) GetAllNodes() ([]model.Node, error) {
	var nodes []model.Node
	if err := nu.nr.GetAllNodes(&nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}
