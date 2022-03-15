package models

type LikeComment struct {
	ID        int
	Type      int
	UserID    int
	CommentID int
}

type LikePost struct {
	ID     int
	Type   int
	UserID int
	PostID int
}
