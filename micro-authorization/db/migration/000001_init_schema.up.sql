CREATE TABLE "role" (
  "id" SMALLSERIAL PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "user_role" (
  "user_id" uuid NOT NULL,
  "role_id" SMALLINT NOT NULL,
  PRIMARY KEY ("user_id", "role_id")
);

CREATE TABLE "permission" (
  "id" SMALLSERIAL PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "role_permission" (
  "role_id" SMALLINT NOT NULL,
  "permission_id" SMALLINT NOT NULL,
  PRIMARY KEY ("role_id", "permission_id")
);

CREATE TABLE "user_banned_permission" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "permission_id" SMALLINT NOT NULL,
  "banned_at" bigint NOT NULL,
  "banned_exp" bigint NOT NULL
);


ALTER TABLE "user_role" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");

ALTER TABLE "user_banned_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");
