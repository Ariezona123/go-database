package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestSekolahInsert(t *testing.T) {
	sekolahRepository := NewSekolahRepository(go_database.GetConnection())

	ctx := context.Background()
	sekolah := entity.Sekolah{
		Guru:  "Killer",
		Mapel: "Matematika",
	}
	result, err := sekolahRepository.Insert(ctx, sekolah)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindByIdSekolah(t *testing.T) {
	sekolahRepository := NewSekolahRepository(go_database.GetConnection())

	sekolah, err := sekolahRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(sekolah)
}

func TestFindAllSekolah(t *testing.T) {
	sekolahRepository := NewSekolahRepository(go_database.GetConnection())

	sekolah, err := sekolahRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, sekolah := range sekolah {
		fmt.Println(sekolah)
	}
}
func TestUpdateSekolah(t *testing.T) {
	sekolahRepository := NewSekolahRepository(go_database.GetConnection())
	ctx := context.Background()
	sekolah := entity.Sekolah{
		Guru:  "Killer",
		Mapel: "Matematika",
	}
	result, err := sekolahRepository.Update(ctx, 1, sekolah)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDeleteSekolah(t *testing.T) {
	sekolahRepository := NewSekolahRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := sekolahRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
