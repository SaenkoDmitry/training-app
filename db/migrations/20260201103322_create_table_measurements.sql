-- +goose Up
-- +goose StatementBegin
CREATE TABLE measurements
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT,
    created_at TIMESTAMP WITH TIME ZONE,
    shoulders  INT,
    chest      INT,
    hands      INT,
    waist      INT,
    buttocks   INT,
    hips       INT,
    calves     INT,
    weight     INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE measurements;
-- +goose StatementEnd
