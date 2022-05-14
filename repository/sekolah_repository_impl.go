package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type sekolahRepositoryImpl struct {
	DB *sql.DB
}

func NewSekolahRepository(db *sql.DB) SekolahRepository {
	return &sekolahRepositoryImpl{db}

}

func (repository *sekolahRepositoryImpl) Insert(ctx context.Context, sekolah entity.Sekolah) (entity.Sekolah, error) {
	script := "INSERT INTO sekolah(guru, mapel) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, sekolah.Guru, sekolah.Mapel)
	if err != nil {
		return sekolah, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return sekolah, err
	}
	sekolah.Id = int32(id)
	return sekolah, nil
	//panic("implement me")
}

func (repository *sekolahRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Sekolah, error) {
	//TODO implement me
	script := "SELECt id, guru, mapel FROM sekolah WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	sekolah := entity.Sekolah{}
	if err != nil {
		return sekolah, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&sekolah.Id, &sekolah.Id, &sekolah.Guru)
		return sekolah, nil
	} else {
		//tidak ada
		return sekolah, errors.New("Id" + strconv.Itoa(int(id)) + " Nor Found")
	}
	//panic("implement me")
}

func (repository *sekolahRepositoryImpl) FindAll(ctx context.Context) ([]entity.Sekolah, error) {
	//TODO implement me
	script := "SELECt id, mapel, guru FROM sekolah"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sekolahs []entity.Sekolah
	for rows.Next() {
		sekolah := entity.Sekolah{}
		rows.Scan(&sekolah.Id, &sekolah.Guru, &sekolah.Mapel)
		sekolahs = append(sekolahs, sekolah)
	}
	return sekolahs, nil
	//panic("implement me")
}

func (repository *sekolahRepositoryImpl) Update(ctx context.Context, id int32, sekolah entity.Sekolah) (entity.Sekolah, error) {
	//TODO implement me
	script := "SELECT id, guru, mapel FROM sekolah WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return sekolah, err
	}
	if rows.Next() {
		// yes
		script := "UPDATE sekolah SET guru = ?, mapel = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, sekolah.Guru, sekolah.Mapel, id)
		if err != nil {
			return sekolah, err
		}
		sekolah.Id = id
		return sekolah, nil
	} else {
		// no
		return sekolah, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *sekolahRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, guru, mapel FROM sekolah WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {

		script := "DELETE FROM sekolah WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {

		return "Gagal", errors.New(("Id" + strconv.Itoa(int(id)) + "tidak ada"))
	}
}
