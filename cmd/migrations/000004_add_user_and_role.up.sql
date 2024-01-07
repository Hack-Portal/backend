

CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" text NOT NULL,
  "role" int NOT NULL
);


CREATE TABLE "roles" (
  "role_id" serial PRIMARY KEY,
  "role" varchar NOT NULL
);

CREATE TABLE "applove_user" (
  "hackathon_id" varchar NOT NULL,
  "user_id" varchar NOT NULL
);

CREATE TABLE "rbac_policies" (
  "p_type" varchar NOT NULL,
  "v0" varchar NOT NULL,
  "v1" varchar NOT NULL,
  "v2" varchar NOT NULL,
  "v3" varchar NOT NULL
);

CREATE INDEX ON "applove_user" ("hackathon_id");

CREATE INDEX ON "applove_user" ("user_id");

CREATE INDEX ON "users" ("user_id");

CREATE INDEX ON "users" ("role");

CREATE INDEX ON "roles" ("role_id");

CREATE INDEX ON "rbac_policies" ("p_type");

CREATE INDEX ON "rbac_policies" ("v0");

CREATE INDEX ON "rbac_policies" ("v1");

CREATE INDEX ON "rbac_policies" ("v2");

CREATE INDEX ON "rbac_policies" ("v3");

ALTER TABLE "users" ADD FOREIGN KEY ("role") REFERENCES "roles" ("role_id");

ALTER TABLE "applove_user" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "applove_user" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
