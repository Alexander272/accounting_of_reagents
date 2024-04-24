-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.notes
(
    id uuid NOT NULL,
    reagent_id uuid NOT NULL,
    comment text COLLATE pg_catalog."default" DEFAULT ''::text,
    note text COLLATE pg_catalog."default" DEFAULT ''::text,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT notes_pkey PRIMARY KEY (id),
    CONSTRAINT notes_reagent_id_fkey FOREIGN KEY (reagent_id)
        REFERENCES public.reagents (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.notes
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.notes;
-- +goose StatementEnd
