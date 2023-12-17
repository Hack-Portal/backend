package models

/*
CREATE TABLE "status_tags" (
  "status_id" serial PRIMARY KEY,
  "status" varchar NOT NULL
);
*/
type StatusTag struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}
