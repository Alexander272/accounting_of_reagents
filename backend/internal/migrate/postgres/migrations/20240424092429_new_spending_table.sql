-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.spending
(
    id uuid NOT NULL,
    reagent_id uuid NOT NULL,
    date_of_spending integer NOT NULL,
    amount real NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT spending_pkey PRIMARY KEY (id),
    CONSTRAINT spending_reagent_id_fkey FOREIGN KEY (reagent_id)
        REFERENCES public.reagents (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.spending
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.spending;
-- +goose StatementEnd
