CREATE TABLE "hackathon_discord_channels" (
  "hackathon_id" varchar NOT NULL,
  "channel_id" varchar NOT NULL
);

CREATE TABLE "discord_server_registries" (
  "channel_id" varchar PRIMARY KEY,
  "channel_name" varchar NOT NULL
);

CREATE TABLE "discord_server_forum_tags" (
  "channel_id" varchar NOT NULL,
  "status_id" int NOT NULL,
  "forum_id" varchar NOT NULL
);

ALTER TABLE "hackathon_discord_channels" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");
ALTER TABLE "discord_server_forum_tags" ADD FOREIGN KEY ("channel_id") REFERENCES "discord_server_registries" ("channel_id");