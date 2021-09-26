package store

import (

	// _ "github.com/lib/pq" //TODO dinamic switch
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	config   *Config
	Database *gorm.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Init and migrations db
func (s *Store) Open() error {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(
		User{},
		UserHistory{},
		UserBid{},
	)
	// if err != nil {
	// 	return err
	// }
	// if err := db.Ping(); err != nil {
	// 	return err
	// }

	s.Database = db

	return nil
}

// func (s *Store) Close() {
// 	s.Database.Close()
// }
