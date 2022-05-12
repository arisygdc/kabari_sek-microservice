CREATE TABLE "auth" (
  "id" uuid PRIMARY KEY,
  "username" VARCHAR,
  "password" VARCHAR
);

CREATE TABLE "auth_user" (
  "auth_id" uuid,
  "user_id" uuid,
  PRIMARY KEY ("auth_id", "user_id")
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "firstname" VARCHAR,
  "lastname" VARCHAR,
  "birth" date
);

CREATE TABLE "friends" (
  "user1" uuid,
  "user2" uuid,
  "approved_at" timestamp,
  PRIMARY KEY ("user1", "user2")
);

ALTER TABLE "auth_user" ADD FOREIGN KEY ("auth_id") REFERENCES "auth" ("id");

ALTER TABLE "auth_user" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "friends" ("user1");

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "friends" ("user2");
