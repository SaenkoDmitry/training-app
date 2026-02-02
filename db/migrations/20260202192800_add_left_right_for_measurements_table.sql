-- +goose Up
-- +goose StatementBegin
ALTER TABLE measurements
    ADD COLUMN hand_left INT;
ALTER TABLE measurements
    ADD COLUMN hand_right INT;
ALTER TABLE measurements
    ADD COLUMN hip_left INT;
ALTER TABLE measurements
    ADD COLUMN hip_right INT;
ALTER TABLE measurements
    ADD COLUMN calf_left INT;
ALTER TABLE measurements
    ADD COLUMN calf_right INT;

UPDATE measurements
SET hand_left  = hands,
    hand_right = hands,
    hip_left   = hips,
    hip_right  = hips,
    calf_left  = calves,
    calf_right = calves
WHERE hands IS NOT NULL
   OR hips IS NOT NULL
   OR calves IS NOT NULL;

ALTER TABLE measurements
DROP
COLUMN hands,
    DROP
COLUMN hips,
    DROP
COLUMN calves;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE measurements
    ADD COLUMN IF NOT EXISTS hands INT;
ALTER TABLE measurements
    ADD COLUMN IF NOT EXISTS hips INT;
ALTER TABLE measurements
    ADD COLUMN IF NOT EXISTS calves INT;

UPDATE measurements
SET hands  = (hand_left + hand_right) / 2,
    hips   = (hip_left + hip_right) / 2,
    calves = (calf_left + calf_right) / 2;

ALTER TABLE measurements DROP COLUMN hand_left;
ALTER TABLE measurements DROP COLUMN hand_right;
ALTER TABLE measurements DROP COLUMN hip_left;
ALTER TABLE measurements DROP COLUMN hip_right;
ALTER TABLE measurements DROP COLUMN calf_left;
ALTER TABLE measurements DROP COLUMN calf_right;
-- +goose StatementEnd
