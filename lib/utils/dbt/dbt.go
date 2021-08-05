package dbt

import (
	"database/sql"
	"wkgc/lib/core/config"
	"wkgc/lib/utils/checkerr"

	_ "github.com/mattn/go-sqlite3"
)

/*
	Dirinfo Model
*/
type DirInfoClass struct {
	Did      int    `col:"did" json:"did"`
	Dirname  string `col:"dirname" json:"dirname"`
	Describe string `col:"describe" json:"describe"`
	Isgit    bool   `col:"isgit" json:"isgit"`
	Lang     string `col:"lang" json:"lang"`
	Tag      string `col:"tag" json:"tag"`
	Created  string `col:"created" json:"created"`
}

/*
	Select Info to Model object
*/
func (dic *DirInfoClass) ParseResult(rows *sql.Rows) {
	for rows.Next() {
		err := rows.Scan(&dic.Did, &dic.Dirname, &dic.Describe, &dic.Isgit, &dic.Lang, &dic.Tag, &dic.Created)
		checkerr.CheckErr(err)
	}
}

/*
	Init LocalDatabase
	Create Table if Table not exists
*/
func InitLocalDatabase() {
	db := OpenLocalDatabase()
	sql_table := `
	CREATE TABLE IF NOT EXISTS "dirinfo" (
		"did" INTEGER PRIMARY KEY AUTOINCREMENT, -- dir id
		"dirname" VARCHAR(64) NULL,
		"describe" VARCHAR(255) NULL,
		"isgit" Bool NULL, -- is git or not
	    "lang" VARCHAR(255) NULL, -- Percentage of each language
	    "tag" VARCHAR(255) NULL, -- project or dir tags
		"created" TIMESTAMP default (datetime('now', 'localtime'))
	);`
	_, err := db.Exec(sql_table)
	checkerr.CheckErr(err)
	db.Close()
}

/*
	Return New sqlite3 sql.DB object
*/
func OpenLocalDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", config.Config.Database)
	checkerr.CheckErr(err)
	return db
}

/*
	Add DirInfo to database
*/
func AddDirInfo(dirname string, describe string, isgit bool) {
	db := OpenLocalDatabase()
	stmt, err := db.Prepare("INSERT INTO dirinfo(dirname, describe, isgit) values(?,?,?)")
	checkerr.CheckErr(err)
	_, err = stmt.Exec(dirname, describe, isgit)
	checkerr.CheckErr(err)
	db.Close()
}

/*
	Select DirInfo from dataase
*/
func SelectDirinfoByDid(did int) *DirInfoClass {
	db := OpenLocalDatabase()
	stmt, err := db.Prepare("SELECT * FROM dirinfo WHERE did=?")
	checkerr.CheckErr(err)
	thisrows, err := stmt.Query(did)
	checkerr.CheckErr(err)
	// 创建对象
	dinfo := new(DirInfoClass)
	// 解析对象
	dinfo.ParseResult(thisrows)
	db.Close()
	return dinfo
}

/*
	Update DirInfo in the database
*/
func UpdateDirInfo(dirinfobj DirInfoClass) {
	db := OpenLocalDatabase()
	stmt, err := db.Prepare("update dirinfo set dirname=?,describe=?,isgit=?,created=?  where did=?")
	checkerr.CheckErr(err)
	_, err = stmt.Exec(dirinfobj.Dirname, dirinfobj.Describe, dirinfobj.Isgit, dirinfobj.Created, dirinfobj.Did)
	checkerr.CheckErr(err)
	db.Close()
}

/*
	Delete DirInfo by did
*/
func DeleteDirInfoByDid(did int) bool {
	db := OpenLocalDatabase()
	stmt, err := db.Prepare("DELETE FROM dirinfo WHERE did=?")
	checkerr.CheckErr(err)
	res, err := stmt.Exec(did)
	checkerr.CheckErr(err)
	affect, err := res.RowsAffected()
	checkerr.CheckErr(err)
	ret := false
	if affect > 0 {
		ret = true
	}
	return ret
}
