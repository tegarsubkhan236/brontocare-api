package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"hospital-api/pkg/api"
	"hospital-api/pkg/repository/model"
	"hospital-api/pkg/repository/seed"
	"log"
	"path/filepath"
	"runtime"
)

type Storage interface {
	RunMigrations(connectionString string, db *sql.DB) error
	RunGormMigrations(gormDB *gorm.DB) error
	api.UserRepository
	api.RoleRepository
	api.AuthRepository
	api.PermissionRepository
	api.HspDisciplineRepository
	api.HspUnitRepository
	api.HspDoctorRepository
}

type storage struct {
	db   *sql.DB
	gorm *gorm.DB
}

func NewStorage(db *sql.DB, gorm *gorm.DB) Storage {
	return &storage{
		db:   db,
		gorm: gorm,
	}
}

func (s *storage) RunMigrations(connectionString string, db *sql.DB) error {
	if connectionString == "" {
		return errors.New("repository: the connString was empty")
	}
	// get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	path := fmt.Sprint(basePath, "/pkg/repository/migrations/")
	migrationsPath := fmt.Sprintf("file:%s", path)
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgresql", driver)

	if err != nil {
		log.Fatal(err)
	}
	// Migrate all the way up ...
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	return nil
}

func (s *storage) RunGormMigrations(gormDB *gorm.DB) error {
	if err := gormDB.AutoMigrate(
		&model.CorePermission{},
		&model.CoreRole{},
		&model.CoreUser{},
		&model.HspDiscipline{},
		&model.HspUnit{},
		&model.HspDoctor{},
	); err != nil {
		return err
	}

	seed.SeedPermission(gormDB)
	seed.SeedRole(gormDB)
	seed.SeedUser(gormDB)

	return nil
}
