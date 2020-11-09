CREATE TABLE "ideas" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);