-- +goose Up
-- +goose StatementBegin
ALTER TABLE rest_timers DROP CONSTRAINT IF EXISTS rest_timers_workout_id_fkey;

ALTER TABLE rest_timers
    ADD CONSTRAINT rest_timers_workout_id_fkey
        FOREIGN KEY (workout_id)
            REFERENCES workout_days (id)
            ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE rest_timers DROP CONSTRAINT IF EXISTS rest_timers_workout_id_fkey;

ALTER TABLE rest_timers
    ADD CONSTRAINT rest_timers_workout_id_fkey
        FOREIGN KEY (workout_id)
            REFERENCES workout_days (id);
-- +goose StatementEnd
