package passwords

type Passwords struct {
	ID       int64  `json:"id" gorm:"primaryKey; unique;  not null"`
	Password string `json:"passwords" gorm:"type:varchar(255);  not null"`
}
