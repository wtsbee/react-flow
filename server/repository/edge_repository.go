package repository

import (
	"react-flow-api/model"

	"gorm.io/gorm"
)

type IEdgeRepository interface {
	GetAllEdges(edge *[]model.Edge) error
}

type edgeRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewEdgeRepository(db *gorm.DB) IEdgeRepository {
	return &edgeRepository{db}
}

func (nr *edgeRepository) GetAllEdges(edges *[]model.Edge) error {
	err := nr.db.Table("edges").Find(&edges).Error
	if err != nil {
		return err
	}
	return nil
}
