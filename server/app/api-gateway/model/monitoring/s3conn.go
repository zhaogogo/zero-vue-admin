package monitoring

type StoreConnectManager struct {
	ID        uint64 `gorm:"column:id"`
	Type      string `gorm:"column:type"`
	Env       string `gorm:"column:env"`
	Host      string `gorm:"column:host"`
	AccessKey string `gorm:"column:access_key"`
	SecretKey string `gorm:"column:secret_key"`
}

func (s *StoreConnectManager) TableName() string {
	return "store_connect_manager"
}
