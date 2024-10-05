# Go CA Expire?

`gocaexpired` is a simple tool to let you know about expired or soon to be expired
certificates from your internal OpenSSL based (sub?) Certificate Authority.

# Building

```shell
go build .
```

# Usage

Just ensure to have a file called `index.txt` in your current working directory.

```shell
./gocaexpired
./gocaexpired -filename /path/to/file/name
```

# Example

```shell
./gocaexpired -filename /etc/ssl/ca/conf/index
Expiry Date: 2022-01-03 CN: git.localnet SN: 819ACDDE2DCEFC8AC7336C4C39502386 CRITICAL: already expired
Expiry Date: 2024-01-03 CN: git.localnet SN: 819ACDDE2DCEFC8AC7336C4C39502387 CRITICAL: already expired
Expiry Date: 2024-10-13 CN: git.localnet SN: 819ACDDE2DCEFC8AC7336C4C39502388 WARNING: expires in < 30d
Expiry Date: 2024-12-05 CN: server.localnet SN: 53DA0418C31D721FC7EE43736780F013 INFO: expires in < 90d
```

# Documentation

* [OpenSSL Index Format](https://pki-tutorial.readthedocs.io/en/latest/cadb.html)
