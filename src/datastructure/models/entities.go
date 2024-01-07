package models

import "time"

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

/*
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
*/
type Hackathon struct {
	HackathonID string    `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int       `json:"term"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

/*
CREATE TABLE "hackathon_status_tags" (
  "hackathon_id" varchar NOT NULL,
  "status_id" int NOT NULL
);
*/

type HackathonStatusTag struct {
	HackathonID string `json:"hackathon_id"`
	StatusID    int64  `json:"status_id"`
}

type JoinedStatusTag struct {
	HackathonID string `json:"hackathon_id"`
	StatusID    int64  `json:"status_id"`
	Status      string `json:"status"`
}

/*
CREATE TABLE "applove_user" (
  "hackathon_id" varchar NOT NULL,
  "user_id" varchar NOT NULL
);
*/

type ApploveUser struct {
	HackathonID string `json:"hackathon_id"`
	UserID      string `json:"user_id"`
}

/*
CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL
);
*/

type User struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
}

/*
CREATE TABLE "roles" (
  "role_id" serial PRIMARY KEY,
  "role" varchar NOT NULL
);
*/

type Role struct {
	RoleID int64  `json:"role_id" gorm:"autoIncrement"`
	Role   string `json:"role"`
}

/*
CREATE TABLE "rbac_policies" (
  "policy_id" int PRIMARY KEY,
  "p_type" varchar NOT NULL,
  "v0" varchar NOT NULL,
  "v1" varchar NOT NULL,
  "v2" varchar NOT NULL,
  "v3" varchar NOT NULL
);
*/

type RbacPolicy struct {
	PolicyID int64  `json:"policy_id" gorm:"autoIncrement"`
	PType    string `json:"p_type"`
	V0       int    `json:"v0"`
	V1       string `json:"v1"`
	V2       string `json:"v2"`
	V3       string `json:"v3"`
}

type CasbinPolicy struct {
	PType string `json:"PType"`
	V0    string `json:"V0"`
	V1    string `json:"V1"`
	V2    string `json:"V2"`
	V3    string `json:"V3"`
}
