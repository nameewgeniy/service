CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "data" json NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);