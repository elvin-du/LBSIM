/*
创建sql语句
*/

package models

import (
	"strings"
)

type Sqler interface {
	GetTableName() string
	GetColumns() []string
	SelectCols() string //需要查询哪些字段
	GetWhere() string
	GetOrder() string //查询出来字段的顺序
	GetLimit() string //查询字段时的限制
}

func InsertSql(sqler Sqler) string {
	columns := sqler.Columns()
	columnStr := "`" + strings.Join(columns, "`,`") + "`"
	placeHolder := strings.Repeat("?,", len(columns))
	sql := fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(%s)", sqler.Tablename(), columnStr, placeHolder[:len(placeHolder)-1])
	return strings.TrimSpace(sql)
}

func SelectSql(sqler Sqler) string {
	where := sqler.GetWhere()
	if where != "" {
		where = "WHERE " + where
	}
	order := sqler.GetOrder()
	if order != "" {
		order = "ORDER BY " + order
	}
	limit := sqler.GetLimit()
	if limit != "" {
		limit = "LIMIT " + limit
	}
	sql := fmt.Sprintf("SELECT %s FROM `%s` %s %s %s", sqler.SelectCols(), sqler.Tablename(), where, order, limit)
	return strings.TrimSpace(sql)
}

func DeleteSql(sqler Sqler) string {
	where := sqler.GetWhere()
	if where != "" {
		where = "WHERE " + where
	}
	sql := fmt.Sprintf("DELETE FROM `%s` %s", sqler.Tablename(), where)
	return strings.TrimSpace(sql)
}

func UpdateSql(sqler Sqler) string {
	columnStr := strings.Join(sqler.Columns(), ",")
	if columnStr == "" {
		return ""
	}
	where := sqler.GetWhere()
	if where != "" {
		where = "WHERE " + where
	}
	sql := fmt.Sprintf("UPDATE `%s` SET %s %s", sqler.Tablename(), columnStr, where)
	return strings.TrimSpace(sql)
}
