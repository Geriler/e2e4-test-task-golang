package migration

import (
	"database/sql"
	"github.com/lopezator/migrator"
)

func Init(db *sql.DB) error {
	m, err := migrator.New(
		migrator.Migrations(
			&migrator.Migration{
				Name: "Create table currencies",
				Func: func(tx *sql.Tx) error {
					if _, err := tx.Exec(`
                        CREATE TABLE IF NOT EXISTS currencies (
                            num_code VARCHAR(3) NOT NULL,
                            char_code VARCHAR(3) NOT NULL,
                            nominal VARCHAR(10) NOT NULL,
                            name VARCHAR(100) NOT NULL,
                            value VARCHAR(100) NOT NULL,
                            date VARCHAR(10) NOT NULL
                        );
					`); err != nil {
						return err
					}
					return nil
				},
			},
		),
	)
	if err != nil {
		return err
	}

	if err = m.Migrate(db); err != nil {
		return err
	}

	return nil
}
