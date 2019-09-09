package migrations

import (
	// stdlib
	"database/sql"
)

// CreateUsersTableUp creates local users table used for authentication.
func CreateUsersTableUp(tx *sql.Tx) error {
	if _, err := tx.Exec(`
	CREATE TABLE users (
		uuid			UUID NOT NULL,
		login			TEXT NOT NULL,
		password_hash	TEXT NOT NULL,
		password_salt	TEXT NOT NULL,
		active			BOOLEAN NOT NULL DEFAULT false,
		created_at		TIMESTAMP WITH TIME ZONE NOT NULL
	);

	COMMENT ON COLUMN users.uuid IS 'User UUID';
	COMMENT ON COLUMN users.login IS 'User login';
	COMMENT ON COLUMN users.password_hash IS 'Hashed user password';
	COMMENT ON COLUMN users.password_salt IS 'Salt for user password';
	COMMENT ON COLUMN users.active IS 'Active user flag. 0 - banned';
	COMMENT ON COLUMN users.created_at IS 'User registration timestamp';

	CREATE INDEX users_uuid_idx ON users(uuid);
	CREATE INDEX users_login_idx ON users(login);
	CREATE INDEX users_created_at_idx ON users(created_at);
	`); err != nil {
		return err
	}

	return nil
}

// CreateUsersTableDown deletes local users table used for authentication.
func CreateUsersTableDown(tx *sql.Tx) error {
	if _, err := tx.Exec(`DROP TABLE users;`); err != nil {
		return err
	}

	return nil
}
