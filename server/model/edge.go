package model

type Edge struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	SourceID uint `json:"source_id"`
	TargetID uint `json:"target_id"`
}

type EdgeDetail struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}
