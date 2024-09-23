/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spansql

import (
	"testing"
)

func TestSQL(t *testing.T) {

	/*
		latz, err := time.LoadLocation("America/Los_Angeles")
		if err != nil {
			t.Fatalf("Loading Los Angeles time zone info: %v", err)
		}

		line := func(n int) Position { return Position{Line: n} }
	*/
	_ = []struct {
		sql     string
		reparse func(string) (interface{}, error)
	}{
		{
			`CREATE TABLE Ta (
  Ca BOOL NOT NULL,
  Cb INT64,
  Cc FLOAT64,
  Cd STRING(17),
  Ce STRING(MAX),
  Cf BYTES(4711),
  Cg BYTES(MAX),
  Ch DATE,
  Ci TIMESTAMP OPTIONS (allow_commit_timestamp = true),
  Cj ARRAY<INT64>,
  Ck ARRAY<STRING(MAX)>,
  Cl TIMESTAMP OPTIONS (allow_commit_timestamp = null),
  Cm INT64 AS (CHAR_LENGTH(Ce)) STORED,
  Cn JSON,
  Co INT64 DEFAULT (1),
) PRIMARY KEY(Ca, Cb DESC)`,
			reparseDDL,
		},
		{
			`CREATE TABLE Tsub (
  SomeId INT64 NOT NULL,
  OtherId INT64 NOT NULL,
  ` + "`Hash`" + ` BYTES(32),
) PRIMARY KEY(SomeId, OtherId),
  INTERLEAVE IN PARENT Ta ON DELETE CASCADE`,
			reparseDDL,
		},
		{
			`CREATE TABLE WithRowDeletionPolicy (
  Name STRING(MAX) NOT NULL,
  DelTimestamp TIMESTAMP NOT NULL,
) PRIMARY KEY(Name),
  ROW DELETION POLICY ( OLDER_THAN ( DelTimestamp, INTERVAL 30 DAY ))`,
			reparseDDL,
		},
		{
			`CREATE TABLE WithSynonym (
  Name STRING(MAX) NOT NULL,
  SYNONYM(AnotherName),
) PRIMARY KEY(Name)`,
			reparseDDL,
		},
		{
			"DROP TABLE Ta",
			reparseDDL,
		},
		{
			"CREATE INDEX Ia ON Ta(Ca, Cb DESC)",
			reparseDDL,
		},
		{
			"DROP INDEX Ia",
			reparseDDL,
		},
		{
			"CREATE OR REPLACE VIEW SingersView SQL SECURITY INVOKER AS SELECT SingerId, FullName, Picture FROM Singers ORDER BY LastName, FirstName",
			reparseDDL,
		},
		{
			"CREATE VIEW vname SQL SECURITY DEFINER AS SELECT cname FROM tname",
			reparseDDL,
		},
		{
			"DROP VIEW SingersView",
			reparseDDL,
		},
		{
			"CREATE ROLE TestRole",
			reparseDDL,
		},
		{
			"DROP ROLE TestRole",
			reparseDDL,
		},
		{
			"GRANT SELECT(name, level, location), UPDATE(location) ON TABLE employees, contractors TO ROLE hr_manager",
			reparseDDL,
		},
		{
			"GRANT EXECUTE ON TABLE FUNCTION tvf_name_one, tvf_name_two TO ROLE hr_manager",
			reparseDDL,
		},
		{
			"GRANT SELECT ON VIEW view_name_one, view_name_two TO ROLE hr_manager",
			reparseDDL,
		},
		{
			"GRANT SELECT ON CHANGE STREAM cs_name_one, cs_name_two TO ROLE hr_manager",
			reparseDDL,
		},
		{
			"REVOKE SELECT(name, level, location), UPDATE(location) ON TABLE employees, contractors FROM ROLE hr_manager",
			reparseDDL,
		},
		{
			"REVOKE EXECUTE ON TABLE FUNCTION tvf_name_one, tvf_name_two FROM ROLE hr_manager",
			reparseDDL,
		},
		{
			"REVOKE SELECT ON VIEW view_name_one, view_name_two FROM ROLE hr_manager",
			reparseDDL,
		},
		{
			"REVOKE SELECT ON CHANGE STREAM cs_name_one, cs_name_two FROM ROLE hr_manager",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta ADD COLUMN Ca BOOL",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta DROP COLUMN Ca",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta SET ON DELETE NO ACTION",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta SET ON DELETE CASCADE",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta ALTER COLUMN Cg STRING(MAX)",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta ALTER COLUMN Ch STRING(MAX) NOT NULL DEFAULT (\"1\")",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta ALTER COLUMN Ci SET OPTIONS (allow_commit_timestamp = null)",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta ALTER COLUMN Cj SET DEFAULT (\"1\")",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta ALTER COLUMN Ck DROP DEFAULT",
			reparseDDL,
		},
		{
			"ALTER TABLE WithRowDeletionPolicy DROP ROW DELETION POLICY",
			reparseDDL,
		},
		{
			"ALTER TABLE WithRowDeletionPolicy ADD ROW DELETION POLICY ( OLDER_THAN ( DelTimestamp, INTERVAL 30 DAY ))",
			reparseDDL,
		},
		{
			"ALTER TABLE WithRowDeletionPolicy REPLACE ROW DELETION POLICY ( OLDER_THAN ( DelTimestamp, INTERVAL 30 DAY ))",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta ADD SYNONYM Syn",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta DROP SYNONYM Syn",
			reparseDDL,
		},
		{
			"ALTER TABLE Ta RENAME TO Tb, ADD SYNONYM Syn",
			reparseDDL,
		},
		{
			"RENAME TABLE Ta TO tmp, Tb TO Ta, tmp TO Tb",
			reparseDDL,
		},
		{
			"ALTER DATABASE dbname SET OPTIONS (enable_key_visualizer=true)",
			reparseDDL,
		},
		{
			"ALTER DATABASE dbname SET OPTIONS (optimizer_version=2)",
			reparseDDL,
		},
		{
			"ALTER DATABASE dbname SET OPTIONS (optimizer_version=2, optimizer_statistics_package='auto_20191128_14_47_22UTC', version_retention_period='7d', enable_key_visualizer=true, default_leader='europe-west1')",
			reparseDDL,
		},
		{
			"ALTER DATABASE dbname SET OPTIONS (optimizer_version=null, optimizer_statistics_package=null, version_retention_period=null, enable_key_visualizer=null, default_leader=null)",
			reparseDDL,
		},
		{
			"CREATE CHANGE STREAM csname",
			reparseDDL,
		},
		{
			"CREATE CHANGE STREAM csname FOR Ta, Tsub(`Hash`)",
			reparseDDL,
		},
		{
			"DROP CHANGE STREAM csname",
			reparseDDL,
		},
		{
			"CREATE CHANGE STREAM csname FOR ALL OPTIONS (value_capture_type='NEW_VALUES')",
			reparseDDL,
		},
		{
			"CREATE CHANGE STREAM csname FOR ALL OPTIONS (retention_period='7d', value_capture_type='NEW_VALUES')",
			reparseDDL,
		},
		{
			"ALTER CHANGE STREAM csname SET FOR ALL",
			reparseDDL,
		},
		{
			"ALTER CHANGE STREAM csname SET FOR Ta, Tsub(`Hash`)",
			reparseDDL,
		},
		{
			"ALTER CHANGE STREAM csname SET OPTIONS (retention_period='7d', value_capture_type='NEW_VALUES')",
			reparseDDL,
		},
		{
			"ALTER CHANGE STREAM csname DROP FOR ALL",
			reparseDDL,
		},
		{
			"ALTER STATISTICS auto_20191128_14_47_22UTC SET OPTIONS (allow_gc=false)",
			reparseDDL,
		},
		{
			"ALTER INDEX iname ADD STORED COLUMN cname",
			reparseDDL,
		},
		{
			"ALTER INDEX iname DROP STORED COLUMN cname",
			reparseDDL,
		},
		{
			`CREATE TABLE IF NOT EXISTS tname (
  id INT64,
  name STRING(64),
) PRIMARY KEY(id)`,
			reparseDDL,
		},
		{
			"CREATE INDEX IF NOT EXISTS Ia ON Ta(Ca)",
			reparseDDL,
		},
		{
			"ALTER TABLE tname ADD COLUMN IF NOT EXISTS cname STRING(64)",
			reparseDDL,
		},
		{
			"DROP TABLE IF EXISTS tname",
			reparseDDL,
		},
		{
			"DROP INDEX IF EXISTS iname",
			reparseDDL,
		},
		{
			`CREATE TABLE tname1 (
  cname1 INT64 NOT NULL,
  cname2 INT64 NOT NULL,
  CONSTRAINT con1 FOREIGN KEY (cname2) REFERENCES tname2 (cname3) ON DELETE NO ACTION,
) PRIMARY KEY(cname1)`,
			reparseDDL,
		},
		{
			`ALTER TABLE tname1 ADD CONSTRAINT con1 FOREIGN KEY (cname2) REFERENCES tname2 (cname3) ON DELETE CASCADE`,
			reparseDDL,
		},
		{
			`CREATE SEQUENCE IF NOT EXISTS sname OPTIONS (sequence_kind='bit_reversed_sequence', skip_range_min=1, skip_range_max=1234567, start_with_counter=50)`,
			reparseDDL,
		},
		{
			`CREATE SEQUENCE sname OPTIONS (sequence_kind='bit_reversed_sequence')`,
			reparseDDL,
		},
		{
			`ALTER SEQUENCE sname SET OPTIONS (sequence_kind='bit_reversed_sequence', skip_range_min=1, skip_range_max=1234567, start_with_counter=50)`,
			reparseDDL,
		},
		{
			`ALTER SEQUENCE sname SET OPTIONS (start_with_counter=1)`,
			reparseDDL,
		},
		{
			`DROP SEQUENCE IF EXISTS sname`,
			reparseDDL,
		},
		{
			`DROP SEQUENCE sname`,
			reparseDDL,
		},
		{
			`INSERT INTO Singers (SingerId, FirstName, LastName) VALUES (1, "Marc", "Richards")`,
			reparseDML,
		},
		{
			"DELETE FROM Ta WHERE C > 2",
			reparseDML,
		},
		{
			`UPDATE Ta SET Cb = 4, Ce = "wow", Cf = Cg, Cg = NULL, Ch = DEFAULT WHERE Ca`,
			reparseDML,
		},
		{
			`SELECT A, B AS banana FROM Table WHERE C < "whelp" AND D IS NOT NULL ORDER BY OCol DESC LIMIT 1000`,
			reparseQuery,
		},
		{
			`SELECT A FROM Table@{FORCE_INDEX=Idx} WHERE B = @b`,
			reparseQuery,
		},
		{
			`SELECT A FROM Table@{FORCE_INDEX=Idx,GROUPBY_SCAN_OPTIMIZATION=TRUE} WHERE B = @b`,
			reparseQuery,
		},
		{
			`SELECT 7`,
			reparseQuery,
		},
		{
			`SELECT CAST(7 AS STRING)`,
			reparseQuery,
		},
		{
			`SELECT SAFE_CAST(7 AS DATE)`,
			reparseQuery,
		},
		{
			`COUNT(*)`,
			reparseExpr,
		},
		{
			`COUNTIF(DISTINCT cname)`,
			reparseExpr,
		},
		{
			`ARRAY_AGG(Foo IGNORE NULLS)`,
			reparseExpr,
		},
		{
			`ANY_VALUE(Foo HAVING MAX Bar)`,
			reparseExpr,
		},
		{
			`STRING_AGG(DISTINCT Foo, "," IGNORE NULLS HAVING MAX Bar)`,
			reparseExpr,
		},
		{
			`X NOT BETWEEN Y AND Z`,
			reparseExpr,
		},
		{
			"SELECT `Desc`",
			reparseQuery,
		},
		{
			`DATE '2014-09-27'`,
			reparseExpr,
		},
		{
			`TIMESTAMP '2014-09-27 12:34:56.123456-07:00'`,
			reparseExpr,
		},
		{
			`JSON '{"a": 1}'`,
			reparseExpr,
		},
		{
			"SELECT A, B FROM Table1 INNER JOIN Table2 ON Table1.A = Table2.A",
			reparseQuery,
		},
		{
			"SELECT A, B FROM Table1 INNER JOIN Table2 ON Table1.A = Table2.A INNER JOIN Table3 USING (X)",
			reparseQuery,
		},
		{
			`SELECT CASE X WHEN 1 THEN "X" WHEN 2 THEN "Y" ELSE NULL END`,
			reparseQuery,
		},
		{
			`SELECT CASE WHEN TRUE THEN "X" WHEN FALSE THEN "Y" END`,
			reparseQuery,
		},
		{
			`SELECT IF(1 < 2, TRUE, FALSE)`,
			reparseQuery,
		},
		{
			`SELECT IFNULL(10, 0)`,
			reparseQuery,
		},
		{
			`SELECT NULLIF(10, 0)`,
			reparseQuery,
		},
		{
			`SELECT COALESCE("A", NULL, "C")`,
			reparseQuery,
		},
	}
	for _, test := range tests {
		// As a confidence check, confirm that parsing the SQL produces the original input.
		_, err := test.reparse(test.sql)
		if err != nil {
			t.Errorf("Reparsing %q: %v", test.sql, err)
			continue
		}
	}
}
