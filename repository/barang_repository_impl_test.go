package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestBarangInsert(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	ctx := context.Background()
	barang := entity.Barang{
		Nama: "Barang Bekas",
		Stok: 220,
	}
	result, err := barangRepository.Insert(ctx, barang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindByIdBarang(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	barang, err := barangRepository.FindById(context.Background(), 2) //jumlh id nya
	if err != nil {
		panic(err)
	}
	fmt.Println(barang)
}

func TestFindAllBarang(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	barang, err := barangRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, barang := range barang {
		fmt.Println(barang)
	}
}
func TestUpdateBarang(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())
	ctx := context.Background()
	barang := entity.Barang{
		Nama: "Barang Luar",
		Stok: 234,
	}
	result, err := barangRepository.Update(ctx, 2, barang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDeleteBarang(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := barangRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
