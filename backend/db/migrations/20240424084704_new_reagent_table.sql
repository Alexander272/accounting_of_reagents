-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.reagents
(
    id uuid NOT NULL,
    type_id uuid NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    uname text COLLATE pg_catalog."default" DEFAULT ''::text,
    document text COLLATE pg_catalog."default" DEFAULT ''::text,
    purity text COLLATE pg_catalog."default" DEFAULT ''::text,
    date_of_manufacture integer DEFAULT 0,
    consignment text COLLATE pg_catalog."default" DEFAULT ''::text,
    manufacturer text COLLATE pg_catalog."default" DEFAULT ''::text,
    shelf_life integer NOT NULL,
    closet text COLLATE pg_catalog."default" NOT NULL,
    shelf integer DEFAULT 0,
    receipt_date integer NOT NULL,
    amount real NOT NULL,
    amount_type_id uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid,
    control_data integer DEFAULT 0,
    protocol text COLLATE pg_catalog."default" DEFAULT ''::text,
    result boolean DEFAULT true,
    seizure text COLLATE pg_catalog."default" DEFAULT ''::text,
    disposal text COLLATE pg_catalog."default" DEFAULT ''::text,
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT reagents_pkey PRIMARY KEY (id),
    CONSTRAINT reagents_type_id_fkey FOREIGN KEY (type_id)
        REFERENCES public.reagent_types (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT reagents_amount_type_id_fkey FOREIGN KEY (amount_type_id)
        REFERENCES public.amount_types (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE SET DEFAULT
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.reagents
    OWNER to postgres;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.reagents;
-- +goose StatementEnd
