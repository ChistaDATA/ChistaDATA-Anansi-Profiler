# ClickHouse-log-analyzer

log format
===========
2022.09.08 05:09:25.696359 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Debug> executeQuery: (from 127.0.0.1:57216)
select * from system.tables; (stage: Complete)
2022.09.08 05:09:25.715277 [ 46 ] {add4a8af-c695-4378-9e55-5345bdea5998} <Information> executeQuery: Read 87 rows,
107.73 KiB in 0.018809208 sec., 4625 rows/sec., 5.59 MiB/sec.

json format (should we flatten it?)
===========

```
{
    "timestamp": "time",
    "queryId": "uuid",
    "initialQueryId": "uuid.UUID",
    "threadIds": [44],
    "client": {
        "host": "string",
        "port": 123,
        "user": "string"
    },
    "query": "string",
    "queryState": {
        "completed": false,
        "error": {
            "code": 1,
            "message": "string",
            "stackTrace": "string"
        }
    },
    "peakMemoryUsage": 12,
    "queryDataInfo": {
        "rowsRead": 1,
        "size": 1,
        "timeElapsed": 1
    }
}
```

output

```
# Query 1: 0 QPS, ID 0x89576B4D8EDF8A30
# Time range: all events occurred at 2020-07-29 16:07:33
# Attribute    pct   total     min     max     avg     95%  stddev  median
# ============ === ======= ======= ======= ======= ======= ======= =======
# Count          0       1
# Exec time     81   1019s   1019s   1019s   1019s   1019s       0   1019s
# Rows examine  98 199.48M 199.48M 199.48M 199.48M 199.48M       0 199.48M
# Bytes sent     0      63      63      63      63      63       0      63
# Query size     0      81      81      81      81      81       0      81
# Peak Memory    1       1       1       1       1       1       0       1 {{Additional}}
# String:
# Databases    suite0db
# Hosts        10.41.0.199
# Users        s0ops
# Query_time distribution
#   1us
#  10us
# 100us
#   1ms
#  10ms
# 100ms
#    1s
#  10s+  ################################################################
# Query
select count(1) from 2873_contact_0 where _73 like '%some' and _74 like '%aaaaa%'\G
```

Reference output
=================

```
# Query 1: 0 QPS, 0x concurrency{{Q1}}, ID 0x89576B4D8EDF8A30 at byte 6934038 __{{Q2}}
# This item is included in the report because it matches --limit.
# Scores: V/M = 0.00{{Q3}}
# Time range: all events occurred at 2020-07-29 16:07:33
# Attribute    pct   total     min     max     avg     95%  stddev  median
# ============ === ======= ======= ======= ======= ======= ======= =======
# Count          0       1
# Exec time     81   1019s   1019s   1019s   1019s   1019s       0   1019s
# Lock time      0    77us    77us    77us    77us    77us       0    77us
# Rows sent      0       1       1       1       1       1       0       1
# Rows examine  98 199.48M 199.48M 199.48M 199.48M 199.48M       0 199.48M
# Rows affecte   0       0       0       0       0       0       0       0
# Bytes sent     0      63      63      63      63      63       0      63
# Query size     0      81      81      81      81      81       0      81
# String:
# Databases    suite0db
# Hosts        10.41.0.199
# Last errno   0
# Users        s0ops
# Query_time distribution
#   1us
#  10us
# 100us
#   1ms
#  10ms
# 100ms
#    1s
#  10s+  ################################################################
# Tables
#    SHOW TABLE STATUS FROM `suite0db` LIKE '2873_contact_0'\G
#    SHOW CREATE TABLE `suite0db`.`2873_contact_0`\G
# EXPLAIN /*!50100 PARTITIONS*/
select count(1) from 2873_contact_0 where _73 like '%some' and _74 like '%aaaaa%'\G
{{Q4}}
```

Info
=====

* R/Call is Average response time per call.

* We can provide an error list
  query id, error code, message, stacktrace

* Add support for cluster. (Aggregation of server logs)

Questions
===========

1. What is 0x concurrency ?

2. WHat is at byte 6934038 __ ?

3. What is V/M ?

4. Qeries ?