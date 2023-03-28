package db

import (
	"RestApi2.0/internal/comment"
	"context"
	"database/sql"
	"fmt"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	id := c.ID
	return comment.Comment{
		Id:     id,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Body.String,
	}
}

func (d *DataBase) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRow("SELECT id, slug,body, author FROM comments WHERE id=$1", uuid)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("Error in fething a comment by id: %w", err)
	}
	return convertCommentRowToComment(cmtRow), nil
}
