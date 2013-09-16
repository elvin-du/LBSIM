/*
对数据库的基本操作
*/

package models

import (
	"config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	*sql.DB
	tableName string
	where     string
	whereVal  []interface{} // where条件对应中字段对应的值
	limit     string
	order     string
	// 插入
	columns   []string      // 需要插入数据的字段
	colValues []interface{} // 需要插入字段对应的值
	// 查询需要
	selectCols string // 想要查询那些字段，接在SELECT之后的，默认为"*"
}

func NewDao(tableName string) *Dao {
	return &Dao{tableName: tableName}
}

func (this *Dao) Open() (err error) {
	this.DB, err = sql.Open(config.Config["driver_name"], config.Config["dsn"])
	return
}

//插入数据
func (this *Dao) Insert() (sql.Result, error) {

}

func (this *Dao) GetColumns() []string {
	return this.columns
}

func (this *Dao) ColValues() []interface{} {
	return this.colValues
}

func (this *Dao) SelectCols() string {
	if this.selectCols == "" {
		return "*"
	}
	return this.selectCols
}

func (this *Dao) GetWhere() string {
	return this.where
}

func (this *Dao) Order(order string) {
	this.order = order
}

func (this *Dao) GetOrder() string {
	return this.order
}

func (this *Dao) Limit(limit string) {
	this.limit = limit
}

func (this *Dao) GetLimit() string {
	return this.limit
}

func (this *Dao) GetTablename() string {
	return this.tablename
}
