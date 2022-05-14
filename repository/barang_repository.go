package repository

import (
	"context"
	"go_database/entity"
)
type BarangRepository interface {
	Insert(ctx context.Context, bola entity.Barang) (entity.Barang, error)
	FindById(ctx context.Context, id int32) (entity.Barang, error)
	FindAll(ctx context.Context) ([]entity.Barang, error)
	Update(ctx context.Context, id int32, barang entity.Barang) (entity.Barang, error)
	Delete(ctx context.Context, Id int32) (string, error)
}