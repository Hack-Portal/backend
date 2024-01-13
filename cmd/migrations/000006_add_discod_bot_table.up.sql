CREATE TABLE "hackathon_discord_channels" (
  "hackathon_id" varchar NOT NULL,
  "channel_id" varchar NOT NULL
);

CREATE TABLE "discord_server_registries" (
  "guild_id" varchar PRIMARY KEY,
  "server_name" varchar NOT NULL,
  "selected_channel" varchar NOT NULL
);

CREATE TABLE "discord_server_forum_tags" (
  "guild_id" varchar NOT NULL,
  "status_id" varchar NOT NULL,
  "forum_id" varchar NOT NULL
);

ALTER TABLE "hackathon_discord_channels" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");
ALTER TABLE "discord_server_forum_tags" ADD FOREIGN KEY ("status_id") REFERENCES "status_tags" ("status_id");