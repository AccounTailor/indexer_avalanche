# Aavalnache Transaction Indexer

This indexer for the Avalanche EVM-compatible C-Chain, a specialized component designed to parse and extract detailed information from individual transactions on the C-Chain. The parsed data is meticulously organized and stored within a NoSQL database to ensure scalability, flexibility, and speed in data retrieval.

This indexer should have the capability to ingest transactions from the Genesis block all the way to the most recent block, storing them in a database. After this initial synchronization, it should continue to monitor and store the transactions from each newly proposed block in real-time into the database.

## Dependencies
1. Install `golang v1.20` or later.
2. Install `MongoDB`.

## How to use indexer?

**MUST** have an environment with a Golang and Mongondibi installed.

Download the dependencies and install the application, run:

```
~$ go install
~$ go get
```

Running the indexer, run:

```
~$ go run main.go
```
