CREATE TABLE "hackathons" (
  "hackathon_id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "icon" text NOT NULL,
  "link" varchar NOT NULL,
  "expired" date NOT NULL,
  "start_date" date NOT NULL,
  "term" int NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "hackathon_status_tags" (
  "hackathon_id" varchar NOT NULL,
  "status_id" int NOT NULL
);

CREATE TABLE "status_tags" (
  "status_id" serial PRIMARY KEY,
  "status" varchar NOT NULL
);

CREATE INDEX ON "hackathons" ("hackathon_id");

CREATE INDEX ON "hackathons" ("expired");

CREATE INDEX ON "hackathons" ("created_at");

CREATE INDEX ON "hackathons" ("start_date");

CREATE INDEX ON "hackathons" ("term");

CREATE INDEX ON "hackathon_status_tags" ("hackathon_id");

CREATE INDEX ON "hackathon_status_tags" ("status_id");

CREATE INDEX ON "status_tags" ("status_id");

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("status_id") REFERENCES "status_tags" ("status_id");
