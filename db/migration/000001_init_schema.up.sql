CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "enteries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "enteries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX "account_by_id" ON "accounts" ("owner");

CREATE INDEX "enteries_by_account_id" ON "enteries" ("account_id");

CREATE INDEX "transfers_by_from" ON "transfers" ("from_account_id");

CREATE INDEX "transfers_by_to" ON "transfers" ("to_account_id");

CREATE INDEX "transfers_by_from_to" ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "enteries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'cannot be negative';
