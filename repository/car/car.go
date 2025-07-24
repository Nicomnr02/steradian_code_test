package carrepository

import (
	"context"
	"steradian_code_test/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryImpl struct {
	DB *pgxpool.Pool
}

type Repository interface {
	Create(c context.Context, data domain.Car) error
	Update(c context.Context, data domain.Car) error
	Delete(c context.Context, ID int) error
	GetAll(c context.Context) ([]domain.Car, error)
	GetByID(c context.Context, ID int) (domain.Car, error)
}

func New(DB *pgxpool.Pool) Repository {
	return &RepositoryImpl{
		DB: DB,
	}
}

func (r *RepositoryImpl) Create(c context.Context, data domain.Car) error {
	sql := `insert into cars (
			car_name,
			day_rate,
			month_rate,
			image
		) values ($1,$2,$3,$4)`

	_, err := r.DB.Exec(c, sql,
		data.CarName,
		data.DayRate,
		data.MonthRate,
		data.Image,
	)

	return err
}
func (r *RepositoryImpl) Update(c context.Context, data domain.Car) error {
	sql := `UPDATE cars SET
		car_name = $1,
		day_rate = $2,
		month_rate = $3,
		image = $4
	WHERE id = $5
	`

	_, err := r.DB.Exec(c, sql,
		data.CarName,
		data.DayRate,
		data.MonthRate,
		data.Image,
		data.ID,
	)
	return err
}

func (r *RepositoryImpl) Delete(c context.Context, ID int) error {
	sql := `DELETE FROM cars WHERE id = $1`
	_, err := r.DB.Exec(c, sql, ID)
	return err
}

func (r *RepositoryImpl) GetAll(c context.Context) ([]domain.Car, error) {
	sql := `
	SELECT
		id,
		car_name,
		day_rate,
		month_rate,
		image
	FROM cars
	`

	rows, err := r.DB.Query(c, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Car])
}

func (r *RepositoryImpl) GetByID(c context.Context, ID int) (domain.Car, error) {
	sql := `SELECT
			id,
			car_name,
			day_rate,
			month_rate,
			image
		FROM cars
		WHERE id = $1`

	rows, err := r.DB.Query(c, sql, ID)
	if err != nil {
		return domain.Car{}, err
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Car])
}
