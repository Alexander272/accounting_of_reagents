-- +goose Up
-- +goose StatementBegin

-- Rename column in role_permissions table
ALTER TABLE public.role_permissions RENAME COLUMN menu_item_id TO permission_id;

ALTER TABLE public.permissions RENAME COLUMN name TO object;
ALTER TABLE public.permissions RENAME COLUMN method TO action;

ALTER TABLE public.permissions ADD COLUMN name TEXT DEFAULT '';

-- Update foreign key if exists
ALTER TABLE IF EXISTS public.role_permissions RENAME CONSTRAINT menu_menu_item_id_fkey TO role_permissions_permission_id_fkey;

ALTER TABLE public.permissions ADD CONSTRAINT permissions_object_action_key UNIQUE (object, action)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE public.role_permissions RENAME COLUMN permission_id TO menu_item_id;

ALTER TABLE IF EXISTS public.role_permissions RENAME CONSTRAINT role_permissions_permission_id_fkey TO menu_menu_item_id_fkey;

ALTER TABLE public.permissions DROP CONSTRAINT permissions_object_action_key;

ALTER TABLE public.permissions DROP COLUMN name;

ALTER TABLE public.permissions RENAME COLUMN object TO name;
ALTER TABLE public.permissions RENAME COLUMN action TO method;

-- +goose StatementEnd
