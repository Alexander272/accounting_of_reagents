-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS public.menu_item RENAME TO permissions;
ALTER TABLE IF EXISTS public.menu RENAME TO role_permissions;

-- Update foreign key constraint names if needed
ALTER INDEX IF EXISTS public.menu_item_pkey RENAME TO permissions_pkey;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS public.permissions RENAME TO menu_item;
ALTER TABLE IF EXISTS public.role_permissions RENAME TO menu;
-- +goose StatementEnd
