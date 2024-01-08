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

CREATE TABLE "proposal_hackathons" (
  "proposal_id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "icon" text NOT NULL,
  "link" varchar NOT NULL,
  "expired" date NOT NULL,
  "start_date" date NOT NULL,
  "term" int NOT NULL,
  "created_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "proposal_hackathon_status_tags" (
  "proposal_id" varchar NOT NULL,
  "status_id" int NOT NULL
);

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "proposal_hackathon_status_tags" ADD FOREIGN KEY ("proposal_id") REFERENCES "proposal_hackathons" ("proposal_id");

ALTER TABLE "proposal_hackathon_status_tags" ADD FOREIGN KEY ("status_id") REFERENCES "status_tags" ("status_id");

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("status_id") REFERENCES "status_tags" ("status_id");