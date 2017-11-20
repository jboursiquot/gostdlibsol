package main

const (
	sqlInsert    = `INSERT INTO proverbs(text, philosopher) VALUES(?,?)`
	sqlSelectAll = `SELECT id, text, philosopher FROM proverbs ORDER BY id`
	sqlSelectOne = `SELECT id, text, philosopher FROM proverbs WHERE id=?`
	sqlUpdate    = `UPDATE proverbs SET text=?, philosopher=? WHERE id=?`
	sqlDelete    = `DELETE FROM proverbs WHERE id=?`
)
