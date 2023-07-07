CREATE TABLE "past_works" (
  "opus" serial PRIMARY KEY,
  "thumbnail_image" bytea NOT NULL,
  "explanatory_text" text NOT NULL
);

CREATE TABLE "hackathons_data" (
  "opus" int NOT NULL,
  "award_id" int NOT NULL,
  "hackathon_id" int NOT NULL
);

CREATE TABLE "awards" (
  "award_id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "past_work_tags" (
  "opus" int NOT NULL,
  "tech_tag_id" int NOT NULL
);

CREATE TABLE "account_tags" (
  "user_id" varchar NOT NULL,
  "tech_tag_id" int NOT NULL
);

CREATE TABLE "tech_tags" (
  "tech_tag_id" serial PRIMARY KEY,
  "tech_tag" varchar NOT NULL
);

CREATE TABLE "accounts" (
  "user_id" varchar PRIMARY KEY,
  "username" int NOT NULL,
  "icon" bytea,
  "explanatory_text" text,
  "locate_id" int NOT NULL,
  "rate" int NOT NULL,
  "hashed_password" varchar,
  "email" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "show_locate" boolean NOT NULL,
  "show_rate" boolean NOT NULL,
  "update_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rate_entries" (
  "user_id" varchar NOT NULL,
  "rate" int NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "locates" (
  "locate_id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "hackathons" (
  "hackathon_id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "icon" bytea,
  "description" text NOT NULL,
  "link" varchar NOT NULL,
  "expired" date NOT NULL,
  "start_date" date NOT NULL,
  "term" int NOT NULL
);

CREATE TABLE "hackathon_status_tags" (
  "hackathon_id" int NOT NULL,
  "status_id" int NOT NULL
);

CREATE TABLE "status_tags" (
  "status_id" serial PRIMARY KEY,
  "status" varchar NOT NULL
);

CREATE TABLE "bookmarks" (
  "hackathon_id" int NOT NULL,
  "user_id" varchar NOT NULL
);

CREATE TABLE "account_past_works" (
  "opus" int NOT NULL,
  "user_id" varchar NOT NULL
);

CREATE TABLE "follows" (
  "to_user_id" varchar NOT NULL,
  "from_user_id" varchar NOT NULL
);

CREATE TABLE "rooms" (
  "room_id" uuid PRIMARY KEY,
  "hackathon_id" int NOT NULL,
  "title" varchar NOT NULL,
  "description" text NOT NULL,
  "limit" int NOT NULL,
  "is_status" boolean NOT NULL
);

CREATE TABLE "rooms_tech_tags" (
  "room_id" uuid NOT NULL,
  "tech_tag_id" int NOT NULL
);

CREATE TABLE "rooms_accounts" (
  "user_id" varchar NOT NULL,
  "room_id" uuid NOT NULL
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("locate_id") REFERENCES "locates" ("locate_id");

ALTER TABLE "account_tags" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("user_id");

ALTER TABLE "account_tags" ADD FOREIGN KEY ("tech_tag_id") REFERENCES "tech_tags" ("tech_tag_id");

ALTER TABLE "past_work_tags" ADD FOREIGN KEY ("tech_tag_id") REFERENCES "tech_tags" ("tech_tag_id");

ALTER TABLE "past_work_tags" ADD FOREIGN KEY ("opus") REFERENCES "past_works" ("opus");

ALTER TABLE "account_past_works" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("user_id");

ALTER TABLE "account_past_works" ADD FOREIGN KEY ("opus") REFERENCES "past_works" ("opus");

ALTER TABLE "rate_entries" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("user_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("to_user_id") REFERENCES "accounts" ("user_id");

ALTER TABLE "follows" ADD FOREIGN KEY ("from_user_id") REFERENCES "accounts" ("user_id");

ALTER TABLE "hackathons_data" ADD FOREIGN KEY ("opus") REFERENCES "past_works" ("opus");

ALTER TABLE "hackathons_data" ADD FOREIGN KEY ("award_id") REFERENCES "awards" ("award_id");

ALTER TABLE "hackathons_data" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "hackathon_status_tags" ADD FOREIGN KEY ("status_id") REFERENCES "status_tags" ("status_id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("user_id");

ALTER TABLE "rooms" ADD FOREIGN KEY ("hackathon_id") REFERENCES "hackathons" ("hackathon_id");

ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("room_id");

ALTER TABLE "rooms_tech_tags" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("room_id");

ALTER TABLE "rooms_tech_tags" ADD FOREIGN KEY ("tech_tag_id") REFERENCES "tech_tags" ("tech_tag_id");

ALTER TABLE "rooms_accounts" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("user_id");


INSERT INTO locates (name) VALUES 
('北海道'), 
('青森県'), 
('岩手県'), 
('宮城県'), 
('秋田県'), 
('山形県'), 
('福島県'), 
('茨城県'), 
('栃木県'), 
('群馬県'), 
('埼玉県'), 
('千葉県'), 
('東京都'), 
('神奈川県'), 
('新潟県'), 
('富山県'), 
('石川県'), 
('福井県'), 
('山梨県'), 
('長野県'), 
('岐阜県'), 
('静岡県'), 
('愛知県'), 
('三重県'), 
('滋賀県'), 
('京都府'), 
('大阪府'), 
('兵庫県'), 
('奈良県'), 
('和歌山県'), 
('鳥取県'), 
('島根県'), 
('岡山県'), 
('広島県'), 
('山口県'), 
('徳島県'), 
('香川県'), 
('愛媛県'), 
('高知県'), 
('福岡県'), 
('佐賀県'), 
('長崎県'), 
('熊本県'), 
('大分県'), 
('宮崎県'), 
('鹿児島県'), 
('沖縄県')
