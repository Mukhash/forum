package db

import (
	"database/sql"
	"fmt"
	"forum/models"
	"unicode"
)

func CreatePost(db *sql.DB, post *models.Post) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	res, err := tx.Exec(
		"INSERT INTO posts (user_id, body, datefrom) VALUES(?, ?, ?)",
		post.UserID, post.Body, post.Datefrom,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	postId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	tags := GetTags(post.Body)
	stmt, err := tx.Prepare("INSERT OR REPLACE INTO tags (id, title) VALUES((SELECT id from tags where tags.title = ?), ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	var tagIds []int
	for _, tag := range tags {
		res, err := stmt.Exec(tag, tag)
		if err != nil {
			tx.Rollback()
			return err
		}
		tagId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		tagIds = append(tagIds, int(tagId))
	}

	stmt2, err := tx.Prepare("INSERT INTO post_tag (post_id, tag_id) VALUES(?,?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt2.Close()

	for _, tagId := range tagIds {
		_, err := stmt2.Exec(postId, tagId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func GetNextPosts(db *sql.DB, firstID, limit int) (*[]models.Post, error) {
	posts := &[]models.Post{}
	queryWhere :=
		`SELECT id, user_id,
		(
			SELECT name FROM users WHERE id = posts.user_id
		) as username, body, datefrom,
		(
			SELECT COUNT(*)
			FROM comments
			WHERE posts.id = comments.post_id
		) AS comments_count,
		(
			SELECT IFNULL (
				SUM (
						CASE WHEN like_type = 1 THEN 1 
							WHEN like_type = 2 THEN -1
						ELSE 0 END
					),
				0)
			FROM post_likes as pl
			WHERE posts.id = pl.post_id
		) AS likes_count,
		IFNULL (
			(
			SELECT IFNULL(like_type, 0)
			FROM post_likes as pl
			WHERE posts.id = pl.post_id AND posts.user_id = pl.user_id
		), 0) AS like_type
		FROM posts
		WHERE id <= ?
		ORDER BY datefrom DESC
		LIMIT ?`
	queryAll :=
		`SELECT id, user_id,
		(
			SELECT name FROM users WHERE id = posts.user_id
		) as username, body, datefrom,
		(
			SELECT COUNT(*)
			FROM comments
			WHERE posts.id = comments.post_id
		) AS comments_count,
		(
			SELECT IFNULL (
				SUM (
						CASE WHEN like_type = 1 THEN 1 
							WHEN like_type = 2 THEN -1
						ELSE 0 END
					),
				0)
			FROM post_likes as pl
			WHERE posts.id = pl.post_id
		) AS likes_count,
		IFNULL (
			(
			SELECT IFNULL(like_type, 0) as like_type
			FROM post_likes as pl
			WHERE posts.id = pl.post_id AND posts.user_id = pl.user_id
		), 0) as like_type
		FROM posts
		ORDER BY datefrom DESC
		LIMIT ?`

	query := queryWhere
	args := []interface{}{firstID, limit}
	if firstID == -1 {
		query = queryAll
		args = []interface{}{limit}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("error rows close: post.go")
		}
	}(rows)

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Username, &post.Body, &post.Datefrom, &post.CommentsCount, &post.LikesCount, &post.LikeType)
		if err != nil {
			return nil, err
		}
		*posts = append(*posts, post)
	}

	return posts, err
}

func GetPost(db *sql.DB, id int) (*models.Post, error) {
	post := &models.Post{}
	if err := db.QueryRow(
		`SELECT	posts.id, posts.user_id,
		(
			SELECT name FROM users WHERE id = posts.user_id
		) as username, posts.body, posts.datefrom,
		(
			SELECT COUNT(*)
			FROM comments
			WHERE posts.id = comments.post_id
		) AS comments_count,
		(
			SELECT IFNULL (
				SUM (
						CASE WHEN like_type = 1 THEN 1 
							WHEN like_type = 2 THEN -1
						ELSE 0 END
					),
				0)
			FROM post_likes as pl
			WHERE posts.id = pl.post_id
		) AS likes_count,
		IFNULL (
			(
			SELECT IFNULL(like_type, 0) as like_type
			FROM post_likes as pl
			WHERE posts.id = pl.post_id AND posts.user_id = pl.user_id
		), 0) as like_type
		FROM posts
		WHERE posts.id = ?`,
		id,
	).Scan(
		&post.ID, &post.UserID, &post.Username, &post.Body, &post.Datefrom, &post.CommentsCount, &post.LikesCount, &post.LikeType,
	); err != nil {
		return nil, err
	}

	return post, nil
}

func GetTags(body string) []string {
	arr := []rune(body)
	var tags []string

	end := 0
	for end < len(arr) {
		if arr[end] == '#' {
			i := end + 1
			for ; i < len(arr); i++ {
				if i == len(arr)-1 && arr[i] != '#' {
					temp := string(arr[end+1 : i+1])
					tags = append(tags, temp)
				}
				if (!unicode.IsLetter(arr[i]) && !unicode.IsDigit(arr[i])) && arr[i] != '#' {
					temp := string(arr[end+1 : i])
					if temp != "" {
						tags = append(tags, temp)
					}
					end = i + 1
					break
				}
				if arr[i] == '#' {
					j := i + 1
					for ; j < len(arr); j++ {
						if !unicode.IsLetter(arr[j]) && !unicode.IsDigit(arr[j]) && arr[j] != '#' {
							j++
							break
						}
						if arr[j] == '#' && j+1 == len(arr) {
							break
						}
						if arr[j] == '#' && arr[j-1] == '#' && arr[j+1] != '#' && !unicode.IsDigit(arr[j+1]) {
							break
						}
					}
					i = j
					break
				}
			}
			end = i
			continue
		}
		end++
	}
	return tags
}

func GetPostsByTag(db *sql.DB, tagID int) ([]models.Post, error) {
	query := `SELECT post_id FROM post_tag
	WHERE tag_id = ?`

	rows, err := db.Query(query, tagID)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var id int
		var post *models.Post

		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		post, err = GetPost(db, id)
		if err != nil {
			return nil, err
		}
		posts = append(posts, *post)
	}
	return posts, nil
}

func GetFavs(db *sql.DB, userID int) ([]models.Post, error) {
	queryPostID := `SELECT post_id FROM post_likes
	WHERE user_id = ? AND like_type = 1`

	rows, err := db.Query(queryPostID, userID)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		var post *models.Post
		post, err = GetPost(db, id)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *post)
	}

	return posts, nil
}

func GetCreated(db *sql.DB, userID int) ([]models.Post, error) {
	query := `SELECT id, user_id,
	(
		SELECT name FROM users WHERE id = posts.user_id
	) as username, body, datefrom,
	(
		SELECT COUNT(*)
		FROM comments
		WHERE posts.id = comments.post_id
	) AS comments_count,
	(
		SELECT IFNULL (
			SUM (
					CASE WHEN like_type = 1 THEN 1 
						WHEN like_type = 2 THEN -1
					ELSE 0 END
				),
			0)
		FROM post_likes as pl
		WHERE posts.id = pl.post_id
	) AS likes_count,
	IFNULL (
		(
		SELECT IFNULL(like_type, 0) as like_type
		FROM post_likes as pl
		WHERE posts.id = pl.post_id AND posts.user_id = pl.user_id
	), 0) as like_type
	FROM posts
	WHERE user_id = ?`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Username, &post.Body, &post.Datefrom, &post.CommentsCount, &post.LikesCount, &post.LikeType)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
