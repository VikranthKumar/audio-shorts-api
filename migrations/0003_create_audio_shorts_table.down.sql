BEGIN;

DROP TRIGGER IF EXISTS audio_shorts_updated_at ON audio_shorts;
DROP TABLE IF EXISTS audio_shorts;
DROP TYPE IF EXISTS audio_shorts_status;
DROP TYPE IF EXISTS category;

COMMIT;