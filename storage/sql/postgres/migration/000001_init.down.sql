ALTER TABLE resources DROP CONSTRAINT "fk_websites";
ALTER TABLE requests DROP CONSTRAINT "fk_websites";
ALTER TABLE requests DROP CONSTRAINT "fk_resources";
ALTER TABLE requests DROP CONSTRAINT "fk_contents";

DROP VIEW urls;

DROP TABLE websites;
DROP TABLE resources;
DROP TABLE contents;
DROP TABLE requests;

DROP EXTENSION hstore;
