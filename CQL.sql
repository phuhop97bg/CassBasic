-- Create a table
CREATE TABLE IF NOT EXISTS dev.users (
userid int PRIMARY KEY,
name text,
age int,
addr text,
last_update_timestamp timestamp
);

-- Insert some data
INSERT INTO dev.users
(userid, name, age, addr, last_update_timestamp)
VALUES (9,'lam an', 23,'nbb', toTimeStamp(now()));

-- Insert Index

CREATE INDEX IF NOT EXISTS age_user
ON dev.users (  age  )

-- Insert custom index
CREATE CUSTOM INDEX stdanalyzer_idx ON dev.users (name) USING 'org.apache.cassandra.index.sasi.SASIIndex'
WITH OPTIONS = {
'mode': 'CONTAINS',
'analyzer_class': 'org.apache.cassandra.index.sasi.analyzer.StandardAnalyzer',
'analyzed': 'true',
'tokenization_skip_stop_words': 'and, the, or',
'tokenization_enable_stemming': 'true',
'tokenization_normalize_lowercase': 'true',
'tokenization_locale': 'en'
};

-- Describe table

SELECT * FROM system_schema.indexes;


DESCRIBE TABLE dev.users;

-- Query
SELECT * FROM dev.users

SELECT * FROM dev.users WHERE age >= 29 ALLOW FILTERING ;

SELECT * FROM dev.users WHERE age >= 29;