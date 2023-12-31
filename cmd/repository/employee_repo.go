package repository

import (
	"errors"
	"service-fleetime/cmd/models"
	"service-fleetime/config"
)

func FetchAllEmployee() (employees []models.Employee, err error) {
	init := config.DB

	query := `
	SELECT * from db_HRMIS.dbo.uf_StrukturOrganisasi();
	`

	if err = init.Raw(query).Scan(&employees).Error; err != nil {
		return employees, err
	}

	return employees, nil
}

func FetchAllEmployeeAndSendToPostgres() (int64, int64, error) {
	init := config.DBPostgres

	employees, err := FetchAllEmployee()

	if err != nil {
		return 0, 0, err
	}

	if len(employees) == 0 {
		return 0, 0, errors.New("data pegawai tidak ditemukan")
	}

	tx := init.Begin()

	result := tx.Exec("DELETE from struktur_organisasi")
	if result.Error != nil {
		tx.Rollback()
		return 0, 0, result.Error
	}
	deletedRows := result.RowsAffected

	queryInsert := `
	INSERT INTO struktur_organisasi (email, kd_kantor, soid, mastersoid, jabatan, sts_jabatan, unitkerja)
	VALUES 
	`
	values := []interface{}{}
	for _, employee := range employees {
		queryInsert += "(?, ?, ?, ?, ?, ?, ?),"
		values = append(
			values,
			employee.EMAIL,
			employee.KD_KANTOR,
			employee.SOID,
			employee.MASTERSOID,
			employee.JABATAN,
			employee.STS_JABATAN,
			employee.UNITKERJA,
		)
	}
	// Remove the last comma
	queryInsert = queryInsert[:len(queryInsert)-1]

	result = tx.Exec(queryInsert, values...)
	if result.Error != nil {
		tx.Rollback()
		return 0, 0, result.Error
	}
	insertedRows := result.RowsAffected

	if err := tx.Commit().Error; err != nil {
		return 0, 0, err
	}

	return insertedRows, deletedRows, nil
}
