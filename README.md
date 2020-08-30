# apd

Append duration to input texts.

# How to install
## by go get
```
# go get github.com/nakatamixi/apd
```
# Usage
```
Usage of apd:
  -f string
    	from key for force specifing
  -p string
    	path to input file
  -t string
    	to key for force specifing
you can specify input from STDIN or -p input file path.
```
# Example

gcloud dataflow jobs durations
```
$ gcloud dataflow jobs list --sort-by="~CREATION_TIME" --filter="name~some_name" --format=json
[
  {
    "creationTime": "2020-08-21 21:30:22",
    "id": "2020-08-21_14_30_19-xxxxxxxxxxx",
    "location": "asia-northeast1",
    "name": "some_name_2020082202",
    "state": "Done",
    "stateTime": "2020-08-21 23:59:07",
    "type": "Batch"
  },
  {
    "creationTime": "2020-08-20 21:30:38",
    "id": "2020-08-20_14_30_36-yyyyyyyyyy",
    "location": "asia-northeast1",
    "name": "some_name_2020082102",
    "state": "Done",
    "stateTime": "2020-08-20 23:55:10",
    "type": "Batch"
  }
]
$ gcloud dataflow jobs list --sort-by="~CREATION_TIME" --filter="name~some_name" --format=json | apd
[
  {
    "creationTime": "2020-08-21 21:30:22",
    "duration": "2h28m45s",
    "id": "2020-08-21_14_30_19-xxxxxxxxxxx",
    "location": "asia-northeast1",
    "name": "some_name_2020082202",
    "state": "Done",
    "stateTime": "2020-08-21 23:59:07",
    "type": "Batch"
  },
  {
    "creationTime": "2020-08-20 21:30:38",
    "duration": "2h24m32s",
    "id": "2020-08-20_14_30_36-yyyyyyyyyy",
    "location": "asia-northeast1",
    "name": "some_name_2020082102",
    "state": "Done",
    "stateTime": "2020-08-20 23:55:10",
    "type": "Batch"
  }
]
```

```
$ mysql -u root test -e "select * from master_xxx"
+----+------+---------------------+---------------------+
| id | name | started_at          | end_at              |
+----+------+---------------------+---------------------+
|  1 | a    | 2020-01-01 00:00:00 | 2020-01-01 00:00:01 |
|  2 | b    | 2020-01-02 00:00:00 | 2020-01-02 01:00:00 |
+----+------+---------------------+---------------------+
$ mysql -u root test -e "select * from master_xxx" |apd
id	name	started_at	end_at	duration
1	a	2020-01-01 00:00:00	2020-01-01 00:00:01	1s
2	b	2020-01-02 00:00:00	2020-01-02 01:00:00	1h0m0s
```

# Disclaimer
- this tool is only for development.
- support format only json, yaml, [c|t]sv
- apd does not keep format such as indent, quotation, line breaking.
