package store

func CreateFriendsTableSQL() {
	sql := "CREATE TABLE IF NOT EXISTS friends (id VARCHAR PRIMARY KEY, dns VARCHAR UNIQUE, name VARCHAR);"
	conn := NewSqliteConn()
	_, err := conn.DB().Exec(sql)

	if err != nil {
		panic(err)
	}
}

func CreateFriendsRequestTableSQL() {
	sql := "CREATE TABLE IF NOT EXISTS friend_requests (id VARCHAR PRIMARY KEY, dns VARCHAR, name VARCHAR, message VARCHAR, status VARCHAR, friend_key VARCHAR, created_at DATETIME DEFAULT CURRENT_TIMESTAMP);"
	conn := NewSqliteConn()
	_, err := conn.DB().Exec(sql)

	if err != nil {
		panic(err)
	}
}

func InitializeDatabase() {
	CreateFriendsTableSQL()
	CreateFriendsRequestTableSQL()
}
