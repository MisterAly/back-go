CREATE TABLE "person" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "document" varchar NOT NULL,
  "phone" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "person_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "person" ("name");

CREATE INDEX ON "account" ("person_id");

COMMENT ON COLUMN "account"."amount" IS 'can be negative';

ALTER TABLE "account" ADD FOREIGN KEY ("person_id") REFERENCES "person" ("id");
