package files

type Files struct {
	ID   int64  `json:"id" gorm:"primaryKey; unique;  not null"`
	Path string `json:"path" gorm:"type:varchar(200); unique;  not null"`
	Type string `json:"type" gorm:"type:varchar(40); not null"`
}
