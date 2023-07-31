package repository

import (
	"react-flow-api/model"

	"gorm.io/gorm"
)

type INodeRepository interface {
	GetAllNodes(node *[]model.Node) error
}

type nodeRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewNodeRepository(db *gorm.DB) INodeRepository {
	return &nodeRepository{db}
}

func (nr *nodeRepository) GetAllNodes(nodes *[]model.Node) error {
	err := nr.db.Table("nodes").Find(&nodes).Error
	if err != nil {
		return err
	}
	return nil
}
