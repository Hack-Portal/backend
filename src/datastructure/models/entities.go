package models

/*
CREATE TABLE "status_tags" (
  "status_id" serial PRIMARY KEY,
  "status" varchar NOT NULL
);
*/
type StatusTag struct {
	// StatusID is primary key and auto increment
	// AutoIncrementなので、Createの際には指定する必要はない
	StatusID int64  `json:"status_id"`
	Status   string `json:"status"`
}
