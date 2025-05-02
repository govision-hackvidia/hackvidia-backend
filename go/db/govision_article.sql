CREATE TABLE "article" (
  "id" int PRIMARY KEY NOT NULL,
  "title" varchar(255) NOT NULL,
  "date" timestampz NOT NULL,
  "source" varchar(255) NOT NULL,
  "read_duration" int NOT NULL,
  "photo" varchar(255) NOT NULL,
  "highlighted_text" varchar(255) NOT NULL,
  "decription" varchar(255) NOT NULL
);
