# HLID: Hexadecimal Lexicographically Sortable Identifier

# Features

HLID strikes a good balance between security, efficiency, and usability:

- 128 bits: fits inside PostgreSQL's UUID type.
- Thanks to the 48 bit timestamp with 100 microsecond resolution, HLIDs are lexicographically sortable. This makes them
  efficient for database indexing. The timestamp won't overflow until 2861-12-16.
- 80 bits of cryptographically secure randomness makes HLIDs usable for secure tokens. You need to generate
  155,885,281,596 HLIDs with the same timestamp to have a 1% chance of collision.
- HLIDs do not contain dashes. This makes HLIDs easier to copy.
- Hexadecimal encoding makes it possible to use HLIDs in PostgreSQL UUID queries without conversions.

## Why not use ...?

- AUTO_INCREMENT/SERIAL
    - Predictable
    - Leaks number of rows
- UUIDv4:
    - Not lexicographically sortable
    - 6 bits "wasted" for version
    - Annoying dashes
- UUIDv7:
    - Lower resolution
    - 6 bits "wasted" for version
    - Annoying dashes
    - [Not cryptographically secure depending on the implementation](https://www.rfc-editor.org/rfc/rfc9562.html#name-monotonicity-and-counters)
- ULID:
    - Lower resolution
    - Text format not compatible with PostgreSQL's UUID type
    - [Not cryptographically secure depending on the implementation](https://github.com/ulid/spec?tab=readme-ov-file#monotonicity)
- KSUID:
    - 20 bytes; not compatible with PostgreSQL's UUID type

## Installation

```
go get github.com/GitRowin/hlid-go
```

## Usage

```go
fmt.Println(hlid.New()) // Example output: 0fca27f9c9315480074ec17d90c3bd52
```
