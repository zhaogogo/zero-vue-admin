package monitoring

type Host struct {
	ID   uint64 `gorm:"primarykey"`
	Host string `gorm:"clumn:host;type:varchar(50);not null;default:''"`
	To   int    `gorm:"clumn:to;type:tinyint;not null"`
	//CreatedAt  uint   `gorm:"autoCreateTime;not null"`
	//ModifiedAt uint   `gorm:"autoUpdateTime;not null"`
	//DeletedAt  uint   `gorm:"index"`

	Tags []HostTag `gorm:"foreignKey:host_id"`
}

func (h *Host) TableName() string {
	return "hosts"
}

type HostTag struct {
	Id     uint64 `gorm:"primaryKey;type:bigint(30) AUTO_INCREMENT"`
	HostId uint64 `gorm:"clumn:host_id"`
	Key    string `gorm:"type:varchar(50);default:'';not null"`
	Value  string `gorm:"type:varchar(50);default:'';not null"`
}

func (t HostTag) TableName() string {
	return "host_tags"
}
