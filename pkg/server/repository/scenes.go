package repository

import (
	"database/sql"
	"prote-API/pkg/server/repository/db"
)

// Scenes Scenesテーブル
type Scenes struct{}

// SelectRowsByName 名前が一致したシーンのSELECT
func (scenes *Scenes) SelectAllRows() ([]ScenesRow, error) {
	rows, err := db.DB.Query("SELECT * FROM scenes")
	if err != nil {
		return nil, err
	}
	return convertRowsToScenesRows(rows)
}

/*

// BulkInsert データベースをレコードを登録する
func (scene *Scene) BulkInsert(name string, num int, works []string, texts []string) error {
	query := "INSERT INTO scenes(name, num, action, text) VALUES"
	queryData := make([]interface{}, num*4, num*4)
	for i := 0; i < num*4; i = i + 4 {
		queryData[i] = name
		queryData[i+1] = i / 4
		queryData[i+2] = works[i/4]
		queryData[i+3] = texts[i/4]
		query += " (?, ?, ?, ?)"
		if i/4 == num-1 {
			break
		} else {
			query += ","
		}
	}
	log.Println(query)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(queryData...)
	return err
}

*/

/*
// Delete データの削除
func (scene *Scene) Delete(name string) error {
	query := "DELETE FROM scene WHERE name = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(name)
	if err != nil {
		return err
	}
	checkNum, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if checkNum == 0 {
		return fmt.Errorf("消すデータが存在しません")
	}
	return err
}
*/

// convertRowsToSceneRows rowsの[]SceneRowへの変換
func convertRowsToScenesRows(rows *sql.Rows) ([]ScenesRow, error) {
	var scenesRows []ScenesRow
	for rows.Next() {
		scenesRow := ScenesRow{}
		err := rows.Scan(&scenesRow.Name, &scenesRow.Num, &scenesRow.Action, &scenesRow.Text, &scenesRow.Image)
		if err != nil {
			return nil, err
		}
		scenesRows = append(scenesRows, scenesRow)
	}
	return scenesRows, nil
}

// ScenesRow scenesテーブルのrow全てを変換するのに使用
type ScenesRow struct {
	Name   string
	Num    int
	Action string
	Text   string
	Image  string
}
