CREATE INDEX ON "hackathons" ("hackathon_id");

CREATE INDEX ON "hackathons" ("expired");

CREATE INDEX ON "hackathons" ("created_at");

CREATE INDEX ON "hackathons" ("start_date");

CREATE INDEX ON "hackathons" ("term");

CREATE INDEX ON "hackathon_status_tags" ("hackathon_id");

CREATE INDEX ON "hackathon_status_tags" ("status_id");

CREATE INDEX ON "status_tags" ("status_id");
