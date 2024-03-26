package controller

import (
	"context"
	"database/sql"
	"log"
	"strconv"
)

//Создать CRUD для таблицы с брэндами
//
//GET /api/v1/brand- получает все брэнды
//GET /api/v1/brand/:id - получает брэнды по id
//POST /api/v1/brand- создает брэнды
//PUT /api/v1/brand/:id - редактирует данные брэнда по id
//DELETE /api/v1/brand/:id - удаляет брэнды по id

type Brand struct {
	BrandId int    `json:"brand_id"`
	Name    string `json:"name"`
	OwnerId int    `json:"owner_id"`
	Phone   string `json:"phone"`
}

func GetAllBrands(ctx context.Context, db *sql.DB) ([]Brand, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM brand")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []Brand
	for rows.Next() {
		var brand Brand

		err := rows.Scan(&brand.BrandId, &brand.Name, &brand.OwnerId, &brand.Phone)
		if err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return brands, nil

}

func GetBrandById(ctx context.Context, db *sql.DB, brandID string) (*Brand, error) {
	row := db.QueryRowContext(ctx, "SELECT * FROM brand where brand_id = $1", brandID)

	var brand Brand
	err := row.Scan(&brand.BrandId, &brand.Name, &brand.OwnerId, &brand.Phone)
	if err != nil {
		return &Brand{}, err
	}
	return &brand, nil
}

func CreateBrand(db *sql.DB, b *Brand) error {
	//var id string

	stmt, err := db.Prepare("INSERT INTO brand(name,owner_id,phone) VALUES ($1,$2,$3) RETURNING brand_id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	//Вставка нового пользователя и получение его айдишника
	//здесь Scan, с помощью которого можно считать все полученные данные в переменные
	//,после запроса возвращаем в id значение
	err = stmt.QueryRow(b.Name, b.OwnerId, b.Phone).Scan(&b.BrandId)

	log.Println("Запрос прошёл")

	if err != nil {
		return err
	}

	return nil
}

// brand_id из запроса
func UpdateBrand(db *sql.DB, b *Brand, brand_id string) (*Brand, error) {
	res, err := db.Exec("UPDATE brand SET name = $1, owner_id = $2, phone = $3 where brand_id = $4", b.Name, b.OwnerId, b.Phone, brand_id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, err
	}

	log.Println("Запрос прошёл")

	b.BrandId, _ = strconv.Atoi(brand_id)

	return b, nil
}

func DeleteBrand(db *sql.DB, brandID string) error {
	_, err := db.Exec("DELETE FROM brand where brand_id = $1", brandID)
	if err != nil {
		return err
	}
	return nil
}
