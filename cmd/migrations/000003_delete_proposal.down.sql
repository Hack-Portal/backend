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