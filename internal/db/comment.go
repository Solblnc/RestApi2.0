package db

import (
	"RestApi2.0/internal/comment"
	"context"
	"database/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		Id:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Body.String,
	}
}

func (d *DataBase) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow

	_, err := d.Client.ExecContext(ctx, "SELECT pg_sleep(16)")
	if err != nil {
		return comment.Comment{}, err
	}
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author 
		FROM commentsv2 
		WHERE id = $1`,
		uuid,
	)
	err = row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error in fething a comment by id: %w", err)
	}
	return convertCommentRowToComment(cmtRow), nil
}

func (d *DataBase) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.Id = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.Id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(ctx,
		`INSERT INTO commentsv2 (id,slug,body, author) VALUES (:id, :slug,:body,:author)`,
		postRow)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	if err = rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil
}

func (d *DataBase) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(ctx, `DELETE from commentsv2 where id=$1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete comment from a database: %w", err)
	}
	return nil
}
func (d *DataBase) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(ctx,
		`UPDATE commentsv2 SET slug=:slug, body=:body, author=:author where id=:id`, cmtRow)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to update comment: %w", err)
	}
	if err = rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}
	return convertCommentRowToComment(cmtRow), nil
}
