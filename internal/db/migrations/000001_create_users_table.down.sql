DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS idx_users_email;
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;