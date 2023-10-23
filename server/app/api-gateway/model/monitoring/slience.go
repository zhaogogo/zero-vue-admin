package monitoring

type SlienceName struct {
	ID            uint64         `gorm:"primarykey"`
	HostID        uint64         `gorm:"clumn:host_id;type:bigint;not null"`
	SlienceName   string         `gorm:"clumn:slience_name;type:varchar(100);not null;default:''"`
	Default       bool           `gorm:"default"`
	SlienceEntrys []SlienceEntry `gorm:"foreignKey:slience_name_id"`
}

func (s *SlienceName) TableName() string {
	return "slience_names"
}

type SlienceEntry struct {
	ID            uint64 `gorm:"primarykey"`
	SlienceNameID uint64 `gorm:"clumn:slience_name_id;type:bigint;not null"`
	Key           string `gorm:"type:varchar(50);default:'';not null"`
	Value         string `gorm:"type:varchar(50);default:'';not null"`
}
