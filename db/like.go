package db

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/models"
)

func InsertLike(db *sql.DB, obj interface{}) error {

	switch obj.(type) {
	case *models.LikePost:
		return handlePost(db, obj)
	case *models.LikeComment:
		return handleComment(db, obj)
	}
	return errors.New("like: Invalid object")
}

func handlePost(db *sql.DB, obj interface{}) error {
	var query string
	var args []interface{}
	likePost := obj.(*models.LikePost)
	//goland:noinspection SqlDialectInspection
	querySelect := `SELECT like_type FROM post_likes
		WHERE user_id = ?
		AND post_id = ?`

	var likeType int
	row := db.QueryRow(querySelect, likePost.UserID, likePost.PostID)
	err := row.Scan(&likeType)

	queryInsert := `INSERT INTO  post_likes
		(like_type, user_id, post_id)
		VALUES (?,?,?)`
	queryDelete := `DELETE FROM post_likes
		WHERE user_id = ?
		AND post_id = ?`
	queryUpdate := `UPDATE post_likes
		SET like_type = ?
		WHERE user_id = ?
		AND post_id = ?`
	fmt.Println(likeType)
	fmt.Println(likePost)
	switch true {
	case err == sql.ErrNoRows:
		query = queryInsert
		args = append(args, likePost.Type, likePost.UserID, likePost.PostID)
	case likeType == likePost.Type:
		query = queryDelete
		args = append(args, likePost.UserID, likePost.PostID)
	case likeType != likePost.Type:
		query = queryUpdate
		args = append(args, likePost.Type, likePost.UserID, likePost.PostID)
	}

	_, err = db.Exec(query, args...)
	return err
}

func handleComment(db *sql.DB, obj interface{}) error {
	var query string
	var args []interface{}
	likeComment := obj.(*models.LikeComment)
	//goland:noinspection SqlDialectInspection
	querySelect := `SELECT like_type FROM comment_likes
		WHERE user_id = ?
		AND comment_id = ?`

	var likeType int
	row := db.QueryRow(querySelect, likeComment.UserID, likeComment.CommentID)
	err := row.Scan(&likeType)

	queryInsert := `INSERT INTO  comment_likes
		(like_type, user_id, comment_id)
		VALUES (?,?,?)`
	queryDelete := `DELETE FROM comment_likes
		WHERE user_id = ?
		AND comment_id = ?`
	queryUpdate := `UPDATE comment_likes
		SET like_type = ?
		WHERE user_id = ?
		AND comment_id = ?`
	fmt.Println(likeType)
	fmt.Println(likeComment)
	switch true {
	case err == sql.ErrNoRows:
		query = queryInsert
		args = append(args, likeComment.Type, likeComment.UserID, likeComment.CommentID)
	case likeType == likeComment.Type:
		query = queryDelete
		args = append(args, likeComment.UserID, likeComment.CommentID)
	case likeType != likeComment.Type:
		query = queryUpdate
		args = append(args, likeComment.Type, likeComment.UserID, likeComment.CommentID)
	}

	_, err = db.Exec(query, args...)
	return err
}
