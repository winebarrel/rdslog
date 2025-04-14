# rdslog

[![CI](https://github.com/winebarrel/rdslog/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/rdslog/actions/workflows/ci.yml)

rdslog is a tool to download RDS logs using the DownloadCompleteDBLogFile API.

see [reference(archive)](https://web.archive.org/web/20171212085731/http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/RESTReference.html).

## Download

https://github.com/kanmu/rdslog/releases/latest

## Usage

```
Usage: rdslog <db-instance-identifier> <log-file-name> [flags]

Arguments:
  <db-instance-identifier>    The customer-assigned name of the DB instance.
  <log-file-name>             The name of the log file.

Flags:
  -h, --help       Show help.
      --version
```

```
$ rdslog my-instance-1 error/postgresql.log.2025-01-23-0456
2025-01-23 04:56:48 UTC::@:[563]:LOG:  starting PostgreSQL 15.12 on x86_64-pc-linux-gnu, compiled by x86_64-pc-linux-gnu-gcc (GCC) 10.5.0, 64-bit
2025-01-23 04:56:48 UTC::@:[563]:LOG:  listening on IPv4 address "0.0.0.0", port 5432
2025-01-23 04:56:48 UTC::@:[563]:LOG:  listening on IPv6 address "::", port 5432
2025-01-23 04:56:48 UTC::@:[563]:LOG:  listening on Unix socket "/tmp/.s.PGSQL.5432"
2025-01-23 04:56:48 UTC::@:[563]:LOG:  Waiting for runtime initialization complete...
2025-01-23 04:56:48 UTC::@:[589]:LOG:  database system was shut down at 2024-06-14 17:27:18 UTC
...
```
