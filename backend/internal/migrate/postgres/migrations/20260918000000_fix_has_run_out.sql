-- +goose Up
-- +goose StatementBegin
-- Развязываем has_run_out от has_notification:
-- ставим has_run_out=true всем реактивам, у которых фактический остаток
-- (amount - сумма списаний) стал <= 0, в том числе уже уведомлённым.
UPDATE public.reagents AS r
SET has_run_out = TRUE
WHERE r.has_run_out = FALSE
  AND r.amount - COALESCE((
        SELECT SUM(s.amount) FROM public.spending AS s WHERE s.reagent_id = r.id
      ), 0) <= 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Откат: снимаем флаг у тех, у кого остаток снова > 0.
UPDATE public.reagents AS r
SET has_run_out = FALSE
WHERE r.has_run_out = TRUE
  AND r.amount - COALESCE((
        SELECT SUM(s.amount) FROM public.spending AS s WHERE s.reagent_id = r.id
      ), 0) > 0;
-- +goose StatementEnd
