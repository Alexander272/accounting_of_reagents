-- +goose Up
-- +goose StatementBegin

-- Remove role_id column and add is_public column to reagent_types
ALTER TABLE reagent_types DROP COLUMN IF EXISTS role_id;
ALTER TABLE reagent_types ADD COLUMN IF NOT EXISTS is_public BOOLEAN DEFAULT FALSE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE reagent_types DROP COLUMN IF EXISTS is_public;
ALTER TABLE reagent_types ADD COLUMN IF NOT EXISTS role_id uuid[];

-- +goose StatementEnd
