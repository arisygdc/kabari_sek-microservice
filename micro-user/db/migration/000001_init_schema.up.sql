CREATE TABLE "auth" (
  "id" uuid PRIMARY KEY,
  "username" VARCHAR NOT NULL,
  "password" VARCHAR NOT NULL,
  "created_at" BIGINT NOT NULL,
  "updated_at" BIGINT NOT NULL
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "firstname" VARCHAR NOT NULL,
  "lastname" VARCHAR NOT NULL,
  "birth" date NOT NULL,
  "created_at" BIGINT NOT NULL,
  "updated_at" BIGINT NOT NULL
);
