CREATE TABLE "movies" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "likes" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "movies" ("id");

CREATE INDEX ON "movies" ("title");
