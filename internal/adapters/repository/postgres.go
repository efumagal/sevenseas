package repository

import (
	"fmt"

	"github.com/efumagal/sevenseas/internal/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PortPostgresRepository struct {
	db *gorm.DB
}

func NewPortPostgresRepository() *PortPostgresRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "pass1234"
	dbname := "postgres"

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.Port{})

	if err != nil {
		panic(err)
	}

	return &PortPostgresRepository{
		db: db,
	}
}

func (m *PortPostgresRepository) SavePort(port domain.Port) error {
	req := m.db.Create(&port)
	if req.RowsAffected == 0 {
		return fmt.Errorf(fmt.Sprintf("port not saved: %v", req.Error))
	}
	return nil
}
