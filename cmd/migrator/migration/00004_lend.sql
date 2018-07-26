-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."lends" (
  "id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  "user_id" text,
  "book_id" uuid,
  "from"   timestamptz DEFAULT now(),
  "to" timestamptz,
  CONSTRAINT "lends_pkey" PRIMARY KEY ("id")
--   CONSTRAINT "book_category_id_fkey" FOREIGN KEY (category_id) REFERENCES categories(id) NOT DEFERRABLE
) WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."lends"