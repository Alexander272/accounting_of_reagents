-- +goose Up
-- +goose StatementBegin
-- Создание дефолтного realm
INSERT INTO realms (id, name, slug, description) 
VALUES ('11111111-0000-0000-0000-000000000001', 'ОИНТО', 'ointo', 'Default realm for existing data')
ON CONFLICT (id) DO NOTHING;

-- Привязка существующих реагентов к дефолтному realm
UPDATE reagents 
SET realm_id = '11111111-0000-0000-0000-000000000001' 
WHERE realm_id IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Отвязка реагентов от дефолтного realm
UPDATE reagents 
SET realm_id = NULL 
WHERE realm_id = '11111111-0000-0000-0000-000000000001';

-- Удаление дефолтного realm
DELETE FROM realms 
WHERE id = '11111111-0000-0000-0000-000000000001';
-- +goose StatementEnd
