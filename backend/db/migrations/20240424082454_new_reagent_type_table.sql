-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.reagent_types
(
    id uuid NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default" DEFAULT ''::text,
    number integer NOT NULL,
    role_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    -- role text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT reagent_types_pkey PRIMARY KEY (id),
    CONSTRAINT reagents_types_role_id_fkey FOREIGN KEY (role_id)
        REFERENCES public.roles (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.reagent_types
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.reagent_types;
-- +goose StatementEnd
