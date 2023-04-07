//go:build integration
// +build integration

package db

import (
	"RestApi2.0/internal/comment"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentDataBase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDataBase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.Id)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)
	})
}
