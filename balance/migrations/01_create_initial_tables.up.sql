CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS balance CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;

SET TIMEZONE="Europe/Moscow";

create table balance (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE,
    currency CHAR(3),
    amount INTEGER CHECK (amount >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

create table transactions (
    id BIGSERIAL PRIMARY KEY,
    transaction_id UUID NOT NULL,
    source VARCHAR(250),
    description VARCHAR(250),
    sender_id UUID REFERENCES balance (user_id) ON DELETE CASCADE NOT NULL,
    recipient_id UUID REFERENCES balance (user_id) ON DELETE CASCADE NOT NULL,
    currency CHAR(3) NOT NULL,
    amount INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX if not exists user_id_idx ON balance (user_id);

CREATE OR REPLACE FUNCTION balance_updated() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER balance_updated_at_trigger
    BEFORE INSERT OR UPDATE
    ON balance
    FOR EACH ROW
EXECUTE PROCEDURE balance_updated();
