BEGIN;

DROP TRIGGER IF EXISTS creators_updated_at ON creators;
DROP TABLE IF EXISTS creators;
DROP TYPE IF EXISTS creator_status;

COMMIT;