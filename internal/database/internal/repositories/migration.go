package repositories

import (
	"database/sql"
	"time"

	"github.com/placeholder/boiler/internal/database/models"
)

func NewMigrationRepository(db *sql.DB) *MigrationRepository {
	return &MigrationRepository{
		conn: db,
	}
}

type MigrationRepository struct {
	conn *sql.DB
}

func (r *MigrationRepository) GetByName(name string) (*models.Migration, error) {
	row := r.conn.QueryRow(`
		SELECT
			id,
			name
		FROM
			migrations
		WHERE
			name = ?
	`, name)
	if row.Err() != nil {
		return nil, row.Err()
	}

	m := &models.Migration{}
	err := row.Scan(&m.Id, &m.Name)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	return m, err
}

func (r *MigrationRepository) ListAll() ([]*models.Migration, error) {
	rows, err := r.conn.Query(`
		SELECT
			id,
			name
		FROM
			migrations
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := []*models.Migration{}

	for rows.Next() {
		m := &models.Migration{}
		if err := rows.Scan(&m.Id, &m.Name); err != nil {
			return nil, err
		}

		results = append(results, m)
	}

	return results, nil
}

func (r *MigrationRepository) Insert(name string) (int64, error) {
	res, err := r.conn.Exec(`
		INSERT INTO migrations (
			name,
			applied_date
		)
		VALUES (
			?,
			?
		)`, name, time.Now().UTC().String())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return id, err
}

func (r *MigrationRepository) DeleteByName(name string) (int64, error) {
	res, err := r.conn.Exec(`
		DELETE FROM
			migrations
		WHERE
			name = ?
		LIMIT 1`, name)
	if err != nil {
		return 0, err
	}

	n, err := res.RowsAffected()
	return n, err
}
