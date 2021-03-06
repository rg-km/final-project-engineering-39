package repository

import (
	"database/sql"
)

type AdminRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{db: db}
}

func (a *AdminRepo) GetTask() ([]Task, error) {
	rows, err := a.db.Query(`
	SELECT
		task.Id,
		task.Judul,
		task.Tanggal,
		penulis.nama AS penulis,
		task.Deskripsi
	FROM task
	INNER JOIN penulis
	ON task.Id_Penulis = penulis.Id`)
	if err != nil {
		return []Task{}, err
	}

	defer rows.Close()

	result := []Task{}
	for rows.Next() {
		admin := Task{}
		err = rows.Scan(&admin.Id, &admin.Judul, &admin.Tanggal, &admin.Penulis, &admin.Deskripsi)
		if err != nil {
			return []Task{}, err
		}
		result = append(result, admin)
	}

	return result, nil
}
func (a *AdminRepo) GetTaskById(id int) (Task, error) {
	row := a.db.QueryRow(`
	SELECT
		task.Id,
		task.Judul,
		task.Tanggal,
		penulis.nama AS penulis,
		task.Deskripsi
	FROM task
	INNER JOIN penulis
	ON task.Id_Penulis = penulis.Id
	WHERE task.Id=?`, id)

	admin := Task{}
	err := row.Scan(&admin.Id, &admin.Judul, &admin.Tanggal, &admin.Penulis, &admin.Deskripsi)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (a *AdminRepo) PutTask(judul string, tanggal string, penulis int, deskripsi string) (int64, error) {
	sqlStatement := `INSERT INTO task (Judul, Tanggal, Id_Penulis, Deskripsi) VALUES (?, ?, ?, ?);`

	stmt, err := a.db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(judul, tanggal, penulis, deskripsi)
	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}
func (a *AdminRepo) UpdateTask(id int, judul string, tanggal string, penulis int, deskripsi string) (int64, error) {
	sqlStatement := `UPDATE task SET Judul = ?, Tanggal = ?, Id_Penulis = ?, Deskripsi = ? WHERE id = ?;`

	stmt, err := a.db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(judul, tanggal, penulis, deskripsi, id)
	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}
func (a *AdminRepo) DeleteTask(id int) (int64, error) {
	sqlStatement := `DELETE FROM task WHERE id = ?;`

	result, err := a.db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	return result.RowsAffected()
}

func (a *AdminRepo) GetPenulis() ([]Penulis, error) {
	rows, err := a.db.Query(`SELECT * FROM penulis;`)
	if err != nil {
		return []Penulis{}, err
	}

	defer rows.Close()

	result := []Penulis{}
	for rows.Next() {
		author := Penulis{}
		err = rows.Scan(&author.Id, &author.Nama)
		if err != nil {
			return []Penulis{}, err
		}
		result = append(result, author)
	}

	return result, nil
}

func (a *AdminRepo) SearchTask(search string) ([]*Task, error) {
	rows, err := a.db.Query("SELECT t.id, t.Judul, t.Tanggal, p.nama, t.Deskripsi FROM task t INNER JOIN penulis p ON t.Id_Penulis = p.id WHERE t.Judul LIKE ?", "%"+search+"%")
	if err != nil {
		return []*Task{}, err
	}

	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Judul, &task.Tanggal, &task.Penulis, &task.Deskripsi)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
