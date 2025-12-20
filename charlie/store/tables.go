package store

import (
	"fmt"
	"os"
)

func ReadSQLFile(prefix string, fileName string) (string, error) {

	sqlFile, err := os.ReadFile(fmt.Sprintf("./store/sql/%s_%s.sql", prefix, fileName))
	if err != nil {
		return "", err
	}

	sql := string(sqlFile)

	return sql, err

}

func CreateTableSQL(tableName string) (string, error) {
	return ReadSQLFile("definition", tableName)
}

func CreateFriendsTableSQL() {
	sql, err := CreateTableSQL("friends")
	if err != nil {
		panic(err)
	}
	conn := NewSqliteConn()
	_, err = conn.DB().Exec(sql)

	if err != nil {
		panic(err)
	}
}

func CreateFriendsRequestTableSQL() {
	sql, err := CreateTableSQL("friend_requests")
	if err != nil {
		panic(err)
	}
	conn := NewSqliteConn()
	_, err = conn.DB().Exec(sql)

	if err != nil {
		panic(err)
	}
}

func CreateSecretsTable() {
}

func InitializeDatabase() {
	CreateFriendsTableSQL()
	CreateFriendsRequestTableSQL()
}
