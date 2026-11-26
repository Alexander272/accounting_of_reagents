-- +goose Up
-- +goose StatementBegin
-- Отчёт по расходу: сначала сужаем списания по партиям и дате,
-- затем фильтруем сами партии по realm/типу.
CREATE INDEX IF NOT EXISTS idx_spending_reagent_date
    ON public.spending (reagent_id, date_of_spending);

CREATE INDEX IF NOT EXISTS idx_reagents_report
    ON public.reagents (realm_id, type_id)
    WHERE deleted IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS public.idx_reagents_report;
DROP INDEX IF EXISTS public.idx_spending_reagent_date;
-- +goose StatementEnd
