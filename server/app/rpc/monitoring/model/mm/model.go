package mm

import (
	"gorm.io/gorm"
	"time"
)

type AlertRules struct {
	ID         uint           `gorm:"primarykey"`
	Ttype      string         `gorm:"clumn:ttype;type:varchar(20);not null;default:''"`
	Name       string         `gorm:"clumn:name;type:varchar(40);not null;default:''"`
	Group      string         `gorm:"clumn:group;type:varchar(3);not null;default:''"`
	To         int            `gorm:"clumn:to;type:tinyint;not null"`
	Expr       string         `gorm:"clumn:expr;type:varchar(255);not null;default:''"`
	Operator   string         `gorm:"clumn:operator;type:varchar(1);not null;default:''"`
	Value      string         `gorm:"clumn:value;type:varchar(20);not null;default:''"`
	For        time.Duration  `gorm:"clumn:for;type:bigint;not null;default:0"`
	AlertText  string         `gorm:"clumn:alert_text;type:varchar(255);not null;default:''"`
	CreatedAt  time.Time      `gorm:"autoCreateTime;not null"`
	ModifiedAt time.Time      `gorm:"autoUpdateTime;not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	Labels []AlertRuleLabel `gorm:"foreignKey:alert_rule_id"`
	Tags   []AlertRuleTag   `gorm:"foreignKey:alert_rule_id"`
}

type AlertRuleLabel struct {
	Id          uint   `gorm:"primaryKey;type:bigint(30) AUTO_INCREMENT"`
	AlertRuleId uint   `gorm:"clumn:alert_rule_id"`
	Key         string `gorm:"type:varchar(50);default:'';not null"`
	Value       string `gorm:"type:varchar(50);default:'';not null"`
}

type AlertRuleTag struct {
	Id          uint   `gorm:"primaryKey;type:bigint(30) AUTO_INCREMENT"`
	AlertRuleId uint   `gorm:"clumn:alert_rule_id"`
	Key         string `gorm:"type:varchar(50);default:'';not null"`
	Value       string `gorm:"type:varchar(50);default:'';not null"`
}
