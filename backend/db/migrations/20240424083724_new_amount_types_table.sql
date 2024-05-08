-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.amount_types
(
    id uuid NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" DEFAULT ''::text,
    number integer NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT amount_types_pkey PRIMARY KEY (id)
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.amount_types
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.amount_types;
-- +goose StatementEnd
