package architectures

import (
	"gorm.io/gorm"
)

type PgDB struct {
	Master *gorm.DB
}
