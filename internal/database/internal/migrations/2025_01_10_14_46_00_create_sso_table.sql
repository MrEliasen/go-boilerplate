CREATE TABLE "sso_connections" (
	"id"	INTEGER,
    "user_id"	INTEGER,
    "nickname"	TEXT,
	"provider_name"	TEXT,
	"provider_userid"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT),
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON UPDATE no action ON DELETE cascade
);
