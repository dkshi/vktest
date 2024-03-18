package repository

import (
	"github.com/dkshi/vktest"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateAdmin(admin *vktest.Admin) (int, error) {
	insertQuery := `INSERT INTO admins (adminname, password) VALUES ($1, $2) RETURNING admin_id;`
	var id int
	res := a.db.QueryRow(insertQuery, admin.Adminname, admin.Password)
	err := res.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetAdmin(adminname, password string) (*vktest.Admin, error) {
	var admin vktest.Admin
	selectQuery := `SELECT adminname, password FROM admins WHERE adminname=$1 AND password=$2`
	res := a.db.QueryRow(selectQuery, adminname, password)
	err := res.Scan(&admin.Adminname, &admin.Password)
	if err != nil {
		return &vktest.Admin{}, err
	}
	return &admin, nil
}
