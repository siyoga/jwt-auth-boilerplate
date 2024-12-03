-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE refresh_tokens
(
	user_id    UUID     NOT NULL,
	number     BIGINT   NOT NULL,
	payload    TEXT 		NOT NULL,
	expires_at BIGINT   NOT NULL
);

ALTER TABLE refresh_tokens
	ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (user_id, number);

ALTER TABLE refresh_tokens
	ADD CONSTRAINT refresh_tokens_user_fkey FOREIGN KEY (user_id) REFERENCES public.users (id);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE refresh_tokens;
