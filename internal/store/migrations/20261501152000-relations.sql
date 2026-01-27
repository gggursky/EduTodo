-- +migrate Up

INSERT INTO courses_vs_themas(course_id,thema_id)
VALUES ((SELECT id FROM courses WHERE code='acoustic'),(SELECT id FROM themas WHERE code='ThemeFrameless'));

INSERT INTO courses_vs_themas(course_id,thema_id)
VALUES ((SELECT id FROM courses WHERE code='acoustic'),(SELECT id FROM themas WHERE code='ThemeAcousticsBasics'));

INSERT INTO courses_vs_themas(course_id,thema_id)
VALUES ((SELECT id FROM courses WHERE code='acoustic'),(SELECT id FROM themas WHERE code='ThemeSoundproofingBasics'));
