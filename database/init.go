package database

import (
	"embed"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // the DB driver
	"github.com/sharpvik/log-go/v2"

	"github.com/lisn-rocks/lisn/configs"
	"github.com/lisn-rocks/lisn/migrations"
)

type Database struct {
	*sqlx.DB
}

func Init() (db *Database) {
	details := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.User,
		configs.Database.Password,
		configs.Database.Name)

	conn := connect(details, 10)

	db = &Database{conn}
	db.up()
	return
}

func connect(details string, tries int) *sqlx.DB {
	if tries < 1 {
		log.Fatal("failed to connect to the database")
	}
	conn, err := sqlx.Connect("postgres", details)
	if err != nil {
		log.Error(err)
		log.Debug("retrying in a second ...")
		time.Sleep(1 * time.Second)
		return connect(details, tries-1)
	}
	return conn
}

func (db *Database) up() {
	log.Debug("applying migrations ...")

	if err := db.applyAll(migrations.Up); err != nil {
		log.Errorf("failed to apply up migrations: %s", err)
		log.Debug("retracting changes ...")
		db.applyAll(migrations.Down)
	}
}

func (db *Database) Down() {
	if err := db.applyAll(migrations.Down); err != nil {
		log.Errorf("failed to apply down migrations: %s", err)
	}
}

func (db *Database) applyAll(fs embed.FS) (err error) {
	migrations, _ := fs.ReadDir(".")
	for _, file := range migrations {
		name := file.Name()
		log.Debug(name)
		migration, _ := fs.ReadFile(name)
		if err := db.apply(migration); err != nil {
			return err
		}
	}
	return
}

func (db *Database) apply(migration []byte) (err error) {
	_, err = db.Exec(string(migration))
	return
}
