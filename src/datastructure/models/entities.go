package models

/*
CREATE TABLE "status_tags" (
  "status_id" serial PRIMARY KEY,
  "status" varchar NOT NULL
);
*/
type StatusTag struct {
	// ID is primary key and auto increment
	// AutoIncrementなので、Createの際には指定する必要はない
	ID     int64  `json:"id"`
	Status string `json:"status"`
}
