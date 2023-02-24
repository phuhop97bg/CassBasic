-- Create a table
CREATE TABLE IF NOT EXISTS keyspace1.users (
userid int PRIMARY KEY,
name text,
age int,
addr text,
last_update_timestamp timestamp
);

-- Insert some data
INSERT INTO keyspace1.users
(userid, name, age, addr, last_update_timestamp)
VALUES (7,'hong bien', 20,'hn', toTimeStamp(now()));

-- Insert Index

CREATE INDEX IF NOT EXISTS age_user
ON keyspace1.users (  age  )

-- Describe table

SELECT * FROM system_schema.indexes;


DESCRIBE TABLE keyspace1.users;

-- Query
SELECT * FROM keyspace1.users

SELECT * FROM keyspace1.users WHERE age >= 29 ALLOW FILTERING ;

SELECT * FROM keyspace1.users WHERE age >= 29;


-- Insert custom index
CREATE CUSTOM INDEX name_idx ON keyspace1.users (name) USING 'org.apache.cassandra.index.sasi.SASIIndex'
WITH OPTIONS = {
'mode': 'CONTAINS',
'analyzed': 'true',
'analyzer_class': 'org.apache.cassandra.index.sasi.analyzer.StandardAnalyzer',
'tokenization_skip_stop_words': 'and, the, or',
'tokenization_enable_stemming': 'true',
'tokenization_normalize_lowercase': 'true',
'tokenization_locale': 'en'
};


CREATE CUSTOM INDEX age_sasi_idx ON keyspace1.users(age) USING 'org.apache.cassandra.index.sasi.SASIIndex' ;



-- Query name

SELECT * FROM keyspace1.users WHERE name LIKE '%hoa%';

SELECT * FROM keyspace1.users WHERE age >= 29 AND age <= 40;
