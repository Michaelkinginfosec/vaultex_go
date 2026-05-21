CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTs accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_id uuid NOT NULL REFERENCES wallets(id) ON DELETE CASCADE,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    account_number VARCHAR(50) NOT NULL,
    account_name VARCHAR(255),
    currency VARCHAR(3) NOT NULL DEFAULT 'NGN',
    ledger_balance DECIMAL(20,2) NOT NULL DEFAULT 0.00,
    locked_balance DECIMAL(20,2) NOT NULL DEFAULT 0.00,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(wallet_id, currency),
    UNIQUE(organization_id, account_number)
);
CREATE INDEX idx_accounts_wallet_id ON accounts(wallet_id);
CREATE INDEX idx_accounts_organization_id ON accounts(organization_id);