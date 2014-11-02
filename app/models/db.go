package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
    DB gorm.DB
)
