DROP TABLE IF EXISTS balance CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TRIGGER IF EXISTS balance_updated_at_trigger ON balance;
DROP FUNCTION IF EXISTS balance_updated;
DROP EXTENSION IF EXISTS "uuid-ossp" CASCADE;