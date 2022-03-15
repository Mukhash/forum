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
			"body"	TEXT,
			"datefrom"	DATETIME,
			FOREIGN KEY ("user_id") REFERENCES users ("id") ON DELETE CASCADE
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
			"user_id"	INTEGER,
			"username" TEXT,
			"body"	TEXT,
			"datefrom"	DATETIME,
			FOREIGN KEY ("post_id") REFERENCES posts ("id") ON DELETE CASCADE
			FOREIGN KEY ("user_id") REFERENCES users ("id") ON DELETE CASCADE
		)`,

		`CREATE TABLE  IF NOT EXISTS "like_types" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"name"	TEXT UNIQUE
		)`,

		`CREATE TABLE  IF NOT EXISTS "post_likes" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"like_type"	INTEGER,
			"user_id"	INTEGER,
			"post_id"	INTEGER,
			FOREIGN KEY ("user_id") REFERENCES users ("id") ON DELETE CASCADE
			FOREIGN KEY ("like_type") REFERENCES like_types ("id")
			FOREIGN KEY ("post_id") REFERENCES posts ("id") ON DELETE CASCADE
			UNIQUE ("like_type","user_id","post_id")
		)`,
		`CREATE TABLE  IF NOT EXISTS "comment_likes" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"like_type"	INTEGER,
			"user_id"	INTEGER,
			"comment_id"	INTEGER,
			FOREIGN KEY ("user_id") REFERENCES users ("id") ON DELETE CASCADE
			FOREIGN KEY ("like_type") REFERENCES like_types ("id")
			FOREIGN KEY ("comment_id") REFERENCES comments ("id") ON DELETE CASCADE
			UNIQUE ("like_type","user_id","comment_id")
		)`,
		`CREATE TABLE IF NOT EXISTS "tags" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"title"	TEXT UNIQUE
		)`,
		`CREATE TABLE  IF NOT EXISTS "post_tag" (
			"id"	INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			"post_id"	INTEGER,
			"tag_id"	INTEGER,
			FOREIGN KEY ("post_id") REFERENCES posts ("id") ON DELETE CASCADE
			FOREIGN KEY ("tag_id") REFERENCES tags ("id") ON DELETE CASCADE
			UNIQUE ("post_id", "tag_id")
		)`,
	}
}
