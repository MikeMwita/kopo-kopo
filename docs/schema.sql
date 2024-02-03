-- schema.sql

-- Create the transactions table
CREATE TABLE IF NOT EXISTS transactions (
                                            transaction_id TEXT PRIMARY KEY,
                                            account_id TEXT NOT NULL,
                                            amount INTEGER NOT NULL,
                                            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create the accounts table
CREATE TABLE IF NOT EXISTS accounts (
                                        account_id TEXT PRIMARY KEY,
                                        balance INTEGER DEFAULT 0
);
