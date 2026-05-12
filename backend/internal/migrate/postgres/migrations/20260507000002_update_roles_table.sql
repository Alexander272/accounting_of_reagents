-- +goose Up
-- +goose StatementBegin

-- Add new columns
ALTER TABLE public.roles ADD COLUMN IF NOT EXISTS slug TEXT NOT NULL DEFAULT '';
ALTER TABLE public.roles ADD COLUMN IF NOT EXISTS is_active BOOLEAN DEFAULT true;
ALTER TABLE public.roles ADD COLUMN IF NOT EXISTS is_system BOOLEAN DEFAULT false;
ALTER TABLE public.roles ADD COLUMN IF NOT EXISTS is_editable BOOLEAN DEFAULT true;
ALTER TABLE public.roles ADD COLUMN IF NOT EXISTS created_at TIMESTAMP WITH TIME ZONE DEFAULT now();
ALTER TABLE public.roles ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP WITH TIME ZONE DEFAULT now();

UPDATE public.roles SET slug = lower(replace(name, ' ', '_')) WHERE slug = '';

-- Drop deprecated is_show column
ALTER TABLE public.roles DROP COLUMN IF EXISTS is_show;

-- Make slug unique
ALTER TABLE public.roles ADD CONSTRAINT roles_slug_unique UNIQUE (slug);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE public.roles DROP CONSTRAINT roles_slug_unique;
ALTER TABLE public.roles DROP COLUMN IF EXISTS slug;
ALTER TABLE public.roles DROP COLUMN IF EXISTS is_active;
ALTER TABLE public.roles DROP COLUMN IF EXISTS is_system;
ALTER TABLE public.roles DROP COLUMN IF EXISTS is_editable;
ALTER TABLE public.roles DROP COLUMN IF EXISTS created_at;
ALTER TABLE public.roles DROP COLUMN IF EXISTS updated_at;

-- Restore is_show column
ALTER TABLE public.roles ADD COLUMN IF NOT EXISTS is_show BOOLEAN DEFAULT true;

-- +goose StatementEnd
