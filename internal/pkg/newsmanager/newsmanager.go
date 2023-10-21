package newsmanager

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Manager struct {
	db *sql.DB
}

type news struct {
	Id    int64
	Title string
	Body  string
	Cover string
}

func (mgr *Manager) Connect(configPath string) error {
	cfg, err := getConfig(configPath)
	if err != nil {
		log.Println("failed to get config")
		return err
	}
	mgr.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	err = mgr.db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Connection with database establised!")
	return nil
}

func (mgr *Manager) GetNews() ([]news, error) {
	rows, err := mgr.db.Query("SELECT * FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []news
	for rows.Next() {
		var n news
		if err := rows.Scan(&n.Id, &n.Title, &n.Body, &n.Cover); err != nil {
			return nil, err
		}
		data = append(data, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func getConfig(configPath string) (*mysql.Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var cfg mysql.Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
