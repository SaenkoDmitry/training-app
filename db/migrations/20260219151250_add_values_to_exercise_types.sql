-- +goose Up
-- +goose StatementBegin
INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Отжимания на брусьях', 'https://disk.yandex.ru/i/lqA5JWOMdY4XPg', 'triceps', 120,
        'грудные мышцы, трицепсы, передние дельтовидные мышцы и мышцы кора', 'reps');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Отжимания от скамьи', 'https://disk.yandex.ru/i/Tw7vxMaGHVc7kA', 'triceps', 120,
        'трицепс и верхняя часть грудных мышц', 'reps');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Разгибание рук в блоке V', 'https://disk.yandex.ru/i/kupoV-btIqWPXA', 'triceps', 120,
        'трехглавая мышца рук', 'reps,weight');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Махи гантелей в стороны', 'https://disk.yandex.ru/i/4Mw6g5m8AJAvmg', 'deltas', 120,
        'передние и задние пучки дельт, вращательная манжета плеча', 'reps,weight');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Протяжка с резинкой стоя', 'https://disk.yandex.ru/i/VBg_r-Bl7S8dww', 'deltas', 120,
        'верхняя часть мускулатуры спины и средних дельт', 'reps,weight');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Армейский жим штанги', 'https://disk.yandex.ru/i/4rxr5d8bg0ka8g', 'deltas', 120,
        'передние дельтовидные мышцы, средние дельты', 'reps,weight');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Махи в стороны в тренажере сидя', 'https://disk.yandex.ru/i/2bFPDljPB_erYA', 'deltas', 120,
        'средние и частично задние пучки дельтовидных мышц', 'reps,weight');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Русские твисты', 'https://disk.yandex.ru/i/yWIT9PBCNTwsDQ', 'press', 90,
        'косые мышцы живота, прямую мышцу живота, поперечную мышцу живота и мышцы нижней части спины', 'reps');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Упражнение ''Велосипед''', 'https://disk.yandex.ru/i/2JOCjtd69ordkA', 'press', 90,
        'прямая, косая и поперечная мышцы живота', 'reps');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Подъем туловища с поднятыми ногами', 'https://disk.yandex.ru/i/YojT_RBFExr4rg', 'press', 90,
        'прямая мышца живота', 'reps');

INSERT INTO exercise_types (name, url, exercise_group_type_code, rest_in_seconds, accent, units)
VALUES ('Скручивания', 'https://disk.yandex.ru/i/ZU4Pley2BXfSVA', 'press', 90,
        'прямая, косая и поперечная мышцы живота', 'reps');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM exercise_types where name = 'Отжимания на брусьях';
DELETE FROM exercise_types where name = 'Отжимания от скамьи';
DELETE FROM exercise_types where name = 'Разгибание рук в блоке V';
DELETE FROM exercise_types where name = 'Махи гантелей в стороны';
DELETE FROM exercise_types where name = 'Протяжка с резинкой стоя';
DELETE FROM exercise_types where name = 'Армейский жим штанги';
DELETE FROM exercise_types where name = 'Махи в стороны в тренажере сидя';
DELETE FROM exercise_types where name = 'Русские твисты';
DELETE FROM exercise_types where name = 'Упражнение ''Велосипед''';
DELETE FROM exercise_types where name = 'Подъем туловища с поднятыми ногами';
DELETE FROM exercise_types where name = 'Скручивания';
-- +goose StatementEnd
