package parsers

import "regexp"

// ClickHouseLogRegEx is the regular expression used to parse clickhouse log
// eg: 2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: (from 127.0.0.1:57216) select * from system.tables; (stage: Complete)
//
//	2022.09.08 05:09:25.698101 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Trace> ContextAccess (default): Access granted: SELECT(database, name, uuid, engine, is_temporary, data_paths, metadata_path, metadata_modification_time, dependencies_database, dependencies_table, create_table_query, engine_full, as_select, partition_key, sorting_key, primary_key, sampling_key, storage_policy, total_rows, total_bytes, lifetime_rows, lifetime_bytes, comment, has_own_data, loading_dependencies_database, loading_dependencies_table, loading_dependent_database, loading_dependent_table) ON system.tables
var ClickHouseLogRegEx *regexp.Regexp

// LogMessageWithQueryInfoRegEx regex for
// eg: executeQuery: (from 127.0.0.1:57236) select * from default.table; (stage: Complete)
//
//	executeQuery: (from 127.0.0.1:53440, initial_query_id: fc1f7dbd-845b-4142-9306-158ddd564e61) INSERT INTO default.data (key) VALUES (stage: Complete)
var LogMessageWithQueryInfoRegEx *regexp.Regexp

// LogMessageWithDataRegEx regex for
// eg: executeQuery: Read 87 rows, 107.73 KiB in 0.018809208 sec., 4625 rows/sec., 5.59 MiB/sec.
var LogMessageWithDataRegEx *regexp.Regexp

// LogMessageWithPeakMemoryRegEx regex for
// eg: MemoryTracker: Peak memory usage (for query): 440.67 KiB.
var LogMessageWithPeakMemoryRegEx *regexp.Regexp

// LogMessageWithErrorRegEx regex for
// eg: executeQuery: Code: 60. DB::Exception: Table default.table doesn't exist. (UNKNOWN_TABLE) (version 22.7.3.5 (official build)) (from 127.0.0.1:57236) (in query: select * from default.table;), Stack trace (when copying this message, always include the lines below):
var LogMessageWithErrorRegEx *regexp.Regexp

// LogMessageWithAccessInfoRegEx regex for
// eg: ContextAccess (default): Access granted: SHOW TABLES ON *.*
var LogMessageWithAccessInfoRegEx *regexp.Regexp

func init() {
	ClickHouseLogRegEx = regexp.MustCompile(`^(\d{4}\.\d{2}\.\d{2} \d{2}:\d{2}:\d{2}\.\d{6}) \[ (\d+) ] \{([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})} <([a-zA-Z]+)> (.*)$`)
	LogMessageWithQueryInfoRegEx = regexp.MustCompile(`^executeQuery: \(from ([a-zA-Z0-9\.:]+):(\d+)(, user: (\w+))?(, initial_query_id: ([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}))?\)(.*)?? (((?i)select|(?i)update|(?i)alter)?.*) \(stage: \w+\)$`)
	LogMessageWithDataRegEx = regexp.MustCompile(`^executeQuery: Read (\d+) rows, (\d+(\.\d+)?) (\w+) in (\d+(\.\d+)?) sec\.,.*$`)
	LogMessageWithPeakMemoryRegEx = regexp.MustCompile(`^MemoryTracker: Peak memory usage( .*)?: (\d+(\.\d+)?) (\w+)\.$`)
	LogMessageWithErrorRegEx = regexp.MustCompile(`^executeQuery: ((Code: (\d+). DB::Exception: (.*))?(.*)??) \(from (.*)\).* \(in query: (.*)\)(, Stack trace \(when copying this message, always include the lines below\)\:)?$`)
	LogMessageWithAccessInfoRegEx = regexp.MustCompile(`^ContextAccess \(\w+\): Access \w+: .* (\w+|\*).(\w+|\*)$`)
}
