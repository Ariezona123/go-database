package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestBolaInsert(t *testing.T) {
	bolaRepository := NewBolaRepository(go_database.GetConnection())

	ctx := context.Background()
	bola := entity.Bola{
		Posisi: "Bek",
		Negara: "Inggris",
	}
	result, err := bolaRepository.Insert(ctx, bola)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindByIdBola(t *testing.T) {
	bolaRepository := NewBolaRepository(go_database.GetConnection())

	bola, err := bolaRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(bola)
}

func TestFindAllBola(t *testing.T) {
	bolaRepository := NewBolaRepository(go_database.GetConnection())

	bola, err := bolaRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, bola := range bola {
		fmt.Println(bola)
	}
}
func TestUpdateBola(t *testing.T) {
	bolaRepository := NewBolaRepository(go_database.GetConnection())
	ctx := context.Background()
	bola := entity.Bola{
		Posisi: "Penyerang",
		Negara: "Indonesia",
	}
	result, err := bolaRepository.Update(ctx, 1, bola)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDeleteBola(t *testing.T) {
	bolaRepository := NewBolaRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := bolaRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
