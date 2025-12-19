CREATE TABLE IF NOT EXISTS coffee_roasters (
    id uuid PRIMARY KEY,
    name text NOT NULL,
    online_shop_url text,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

