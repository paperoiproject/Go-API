package repository

import (
	"database/sql"
	"log"
	"prote-API/pkg/server/repository/db"
	"time"
)

type ScenarioTable struct{}

// SelectAll 名前が一致したシーンのSELECT
func (scenarioTable *ScenarioTable) SelectAll() ([]ScenarioRow, error) {
	rows, err := db.DB.Query("SELECT * FROM scenario")
	if err != nil {
		return nil, err
	}
	return convertRowsToScenarioRows(rows)
}

// Insert シナリオの追加
func (scenarioTable *ScenarioTable) Insert(name string, date time.Time) error {
	stmt, err := db.DB.Prepare("INSERT INTO scenario(name, date) VALUES (?, ?)")
	_, err = stmt.Exec(name, date)
	return err
}

// Make シナリオの作成
func (scenarioTable *ScenarioTable) Make(name string, date time.Time, acts []string, texts []string, images []string) error {
	log.Println("hello", name)
	tx, err := db.DB.Begin()
	if err != nil {
		log.Println("erro")
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			log.Println("recover")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println("rollback")
			tx.Rollback()
		} else {
			log.Println("commit")
			err = tx.Commit()
		}
	}()
	// シナリオテーブルへの追加
	stmt, err := tx.Prepare("INSERT INTO scenario(name, date) VALUES (?, ?)")
	log.Println("aaa", date)
	_, err = stmt.Exec(name, date)
	if err != nil {
		return err
	}
	log.Println("OK: scenario")
	// <---
	// シーンズテーブルへの追加
	query := "INSERT INTO scenes(name, num, action, text, image) VALUES"
	n := len(acts)
	queryData := make([]interface{}, n*5, n*5)
	for i := 0; i < n*5; i = i + 5 {
		queryData[i] = name
		queryData[i+1] = i / 5
		queryData[i+2] = acts[i/5]
		queryData[i+3] = texts[i/5]
		queryData[i+4] = images[i/5]
		query += " (?, ?, ?, ?, ?)"
		if i+5 == n*5 {
			break
		} else {
			query += ","
		}
	}
	log.Println(query)
	stmt, err = tx.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(queryData...)
	return err
}

func convertRowsToScenarioRows(rows *sql.Rows) ([]ScenarioRow, error) {
	var scenarioRows []ScenarioRow
	for rows.Next() {
		scenarioRow := ScenarioRow{}
		err := rows.Scan(&scenarioRow.Name, &scenarioRow.Date)
		if err != nil {
			return nil, err
		}
		scenarioRows = append(scenarioRows, scenarioRow)
	}
	return scenarioRows, nil
}

type ScenarioRow struct {
	Name string
	Date time.Time
}
