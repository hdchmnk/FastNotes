-- +goose Up
-- +goose StatementBegin
CREATE TABLE notes
(
    id SERIAL PRIMARY KEY,
    userID SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(MAX) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notes;
-- +goose StatementEnd
