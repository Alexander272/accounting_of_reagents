-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.reagents
    ADD COLUMN realm_id uuid;

ALTER TABLE public.reagents
    ADD CONSTRAINT reagents_realm_id_fkey FOREIGN KEY (realm_id)
        REFERENCES public.realms (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.reagents
    DROP CONSTRAINT IF EXISTS reagents_realm_id_fkey;

ALTER TABLE public.reagents
    DROP COLUMN IF EXISTS realm_id;
-- +goose StatementEnd
