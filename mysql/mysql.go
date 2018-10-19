package mysql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

var (
	// CommentTag 注释标签
	CommentTag = "comment"
	// CommentDelim 注释分隔符
	CommentDelim = ","
)

// Table 数据表
type Table interface {
	// 表名
	TableName() string
	// 表注释
	TableComment() string
}

// Logger 日志接口
type Logger interface {
	Printf(format string, args ...interface{})
}

// NewDB 创建数据库实例
func NewDB(db *sql.DB, schema string, logger ...Logger) *DB {
	d := &DB{
		db:     db,
		schema: schema,
	}
	if len(logger) > 0 {
		d.out = logger[0]
	}
	return d
}

// DB 数据库管理
type DB struct {
	schema string
	db     *sql.DB
	out    Logger
}

func (d *DB) debug(format string, args ...interface{}) {
	if d.out != nil {
		d.out.Printf(format, args...)
	}
}

// AlterComment 更新注释(表注释及列注释)
func (d *DB) AlterComment(tables ...Table) error {
	for _, t := range tables {
		if v := t.TableComment(); v != "" {
			err := d.alterTableComment(t.TableName(), v)
			if err != nil {
				return err
			}
		}

		err := d.alterColumnComment(t)
		if err != nil {
			return err
		}

	}

	return nil
}

// 更新表注释
func (d *DB) alterTableComment(name, comment string) error {
	query := fmt.Sprintf("ALTER TABLE %s COMMENT '%s'", name, comment)
	d.debug(query)
	_, err := d.db.Exec(query)
	return err
}

func (d *DB) alterColumnComment(table Table) error {
	iVal := reflect.Indirect(reflect.ValueOf(table))
	if iVal.Kind() != reflect.Struct {
		return fmt.Errorf("invalid table")
	}

	for i := 0; i < iVal.NumField(); i++ {
		field := iVal.Type().Field(i)
		tag := field.Tag.Get(CommentTag)
		if tag == "" {
			continue
		}

		var (
			column  string
			comment string
		)
		ss := strings.Split(tag, CommentDelim)
		if len(ss) == 2 {
			column = ss[0]
			comment = ss[1]
		} else {
			comment = ss[0]
			column = field.Name
		}

		columnType, err := d.getColumnType(table.TableName(), column)
		if err != nil {
			return err
		}

		query := fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s %s COMMENT '%s';", table.TableName(), column, columnType, comment)
		d.debug(query)
		_, err = d.db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DB) getColumnType(table, column string) (string, error) {
	query := fmt.Sprintf("SELECT column_type FROM information_schema.columns WHERE table_schema ='%s' AND table_name = '%s' AND column_name='%s'; ", d.schema, table, column)
	d.debug(query)

	var columnType string
	row := d.db.QueryRow(query)
	err := row.Scan(&columnType)
	return columnType, err
}
