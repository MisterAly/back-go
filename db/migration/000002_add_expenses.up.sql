CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "value" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "expenses" ("account_id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");
