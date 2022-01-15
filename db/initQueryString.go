package db

func getQuery() []string {
	return []string{
		`CREATE TABLE  IF NOT EXISTS "users" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"name"	TEXT UNIQUE,
			"password"	TEXT,
			"email"	TEXT UNIQUE
		)`,

		`CREATE TABLE  IF NOT EXISTS "auth_sessions" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"user_id"	INTEGER,
			"cookie_value"	TEXT UNIQUE,
			"status"	INTEGER DEFAULT 0,
			"dateto"	DATETIME,
			FOREIGN KEY ("user_id") REFERENCES users ("id") ON DELETE CASCADE
		)`,

		`CREATE TABLE  IF NOT EXISTS "posts" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"user_id"	INTEGER,
			"text"	TEXT,
			"datefrom"	DATETIME,
			FOREIGN KEY ("user_id") REFERENCES users ("id") ON DELETE CASCADE
		)`,

		`CREATE TABLE	IF NOT EXISTS "categories" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"name"	TEXT UNIQUE
		)`,

		`CREATE TABLE  IF NOT EXISTS "post_category" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"post_id"	INTEGER,
			"category_id"	INTEGER,
			FOREIGN KEY ("post_id") REFERENCES posts ("id") ON DELETE CASCADE
			FOREIGN KEY ("category_id") REFERENCES categories ("id") ON DELETE CASCADE
			UNIQUE("post_id", "category_id")
		)`,

		`CREATE TABLE  IF NOT EXISTS "comments" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"post_id"	INTEGER,
			"uid"	INTEGER,
			"text"	TEXT,
			"creation_date"	DATETIME,
			FOREIGN KEY ("post_id") REFERENCES posts ("id") ON DELETE CASCADE
			FOREIGN KEY ("uid") REFERENCES users ("id") ON DELETE CASCADE
		)`,

		`CREATE TABLE  IF NOT EXISTS "rate_type" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"name"	TEXT UNIQUE
		)`,

		`CREATE TABLE IF NOT EXISTS "obj_type" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"name"	TEXT UNIQUE
		)`,

		`CREATE TABLE  IF NOT EXISTS "rates" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"rate_type"	INTEGER,
			"obj_type"	INTEGER,
			"uid"	INTEGER,
			"obj_id"	INTEGER,
			FOREIGN KEY ("uid") REFERENCES users ("id") ON DELETE CASCADE
			FOREIGN KEY ("rate_type") REFERENCES rate_type ("id")
			FOREIGN KEY ("obj_type") REFERENCES obj_type ("id")
			UNIQUE ("rate_type","obj_type","uid","obj_id")
		)`,
	}
}
