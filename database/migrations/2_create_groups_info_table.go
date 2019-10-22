package migrations

import (
	// stdlib
	"database/sql"
)

// CreateGroupsTableUp creates table with groups information.
func CreateGroupsTableUp(tx *sql.Tx) error {
	if _, err := tx.Exec(`
	CREATE TABLE groups (
		uuid				UUID NOT NULL,
		group_name			TEXT NOT NULL,
		group_description	TEXT,
		articles_count		INTEGER NOT NULL DEFAULT 0,
		created_at			TIMESTAMP WITH TIME ZONE NOT NULL,
		last_message_at		TIMESTAMP WITH TIME ZONE NOT NULL
	);

	COMMENT ON COLUMN groups.uuid IS 'Group UUID';
	COMMENT ON COLUMN groups.group_name IS 'Group name';
	COMMENT ON COLUMN groups.group_description IS 'Group description';
	COMMENT ON COLUMN groups.articles_count IS 'Articles count in group';
	COMMENT ON COLUMN groups.created_at IS 'Group creation timestamp';
	COMMENT ON COLUMN groups.last_message_at IS 'When last message appeared in this group?';

	CREATE INDEX groups_uuid_idx ON groups(uuid);
	CREATE INDEX groups_name_idx ON groups(group_name);
	CREATE INDEX groups_created_at_idx ON groups(created_at);

	`); err != nil {
		return err
	}

	return nil
}

// CreateGroupsTableDown deletes table with groups information.
func CreateGroupsTableDown(tx *sql.Tx) error {
	if _, err := tx.Exec(`DROP TABLE groups;`); err != nil {
		return err
	}

	return nil
}
