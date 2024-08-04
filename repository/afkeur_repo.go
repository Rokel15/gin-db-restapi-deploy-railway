package repository

import (
	"database/sql"
	"golang-gin-db-restapi-lokal/entities"
)

func GetAllAfkeur(db *sql.DB) (result []entities.Afkeur, err error) {
	sql := "SELECT * FROM afkeur"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var data entities.Afkeur
		err = rows.Scan(&data.ID, &data.JumlahAyam, &data.StartDate, &data.EndDate, &data.Mortalitas)
		if err != nil {
			return
		}
		result = append(result, data)
	}
	return
}

func InserAfkeur(db *sql.DB, afkeur entities.Afkeur) (err error) {
	sql := "INSERT INTO afkeur(ID, jumlahAyam, startDate, endDate, Mortalitas) values($1, $2, $3, $4, $5)"

	errs := db.QueryRow(sql, afkeur.ID, afkeur.JumlahAyam, afkeur.StartDate, afkeur.EndDate, afkeur.Mortalitas)

	return errs.Err()
}

func UpdateAfkeur(db *sql.DB, afkeur entities.Afkeur) (err error) {
	sql := "UPDATE afkeur SET jumlahAyam = $1, startDate = $2, endDate = $3 WHERE ID = $4"

	// errs := db.QueryRow(sql, afkeur.JumlahAyam, afkeur.StartDate, afkeur.EndDate, afkeur.ID)
	_, err = db.Exec(sql, afkeur.JumlahAyam, afkeur.StartDate, afkeur.EndDate, afkeur.ID)

	// return errs.Err()
	return
}

func DeleteAfkeur(db *sql.DB, afkeur entities.Afkeur) (err error) {
	sql := "DELETE FROM afkeur WHERE ID = $1"

	// errs := db.QueryRow(sql, afkeur.ID)
	_, err = db.Exec(sql, afkeur.ID)

	// return errs.Err()
	return
}
