CREATE TABLE "hackathon_proposals" (
  "hackathon_proposal_id" serial PRIMARY KEY,
  "url" varchar NOT NULL,
  "is_approved" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" timestamptz NOT NULL
);