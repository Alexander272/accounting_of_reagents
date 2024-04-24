-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.extending
(
    id uuid NOT NULL,
    reagent_id uuid NOT NULL,
    date_of_extending integer NOT NULL,
    period_of_extending integer NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT extending_pkey PRIMARY KEY (id),
    CONSTRAINT extending_reagent_id_fkey FOREIGN KEY (reagent_id)
        REFERENCES public.reagents (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.extending
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.extending;
-- +goose StatementEnd
