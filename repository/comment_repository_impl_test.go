package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "arieafriyanto@gmail.com",
		Comment: "Kamu Pasti Bisa",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestUpdateComment(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "arieafriyanto18@gmail.com",
		Comment: "Rajin belajar ya",
	}
	result, err := commentRepository.Update(ctx, 1, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDeleteComment(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := commentRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
