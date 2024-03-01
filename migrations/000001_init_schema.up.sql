CREATE TABLE IF NOT EXISTS "accounts" (
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "username" text NOT NULL,
    "balance" bigint NOT NULL DEFAULT 0,
    "currency" text NOT NULL
);

CREATE TABLE IF NOT EXISTS "entries" (
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "account_id" bigint NOT NULL,
    "amount" bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS "transfers" (
    "id" bigserial PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "from_account_id" bigint NOT NULL,
    "to_account_id" bigint NOT NULL,
    "amount" bigint NOT NULL
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("username");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
