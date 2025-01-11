CREATE TABLE "migrations" (
	"id"	INTEGER,
	"name"	TEXT UNIQUE,
	"applied_date"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
