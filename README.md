# Audio Shorts API

A simple GraphQL Audio Shorts API to perform simple CRUD of audio shorts, built to showcase a backend service written in Go 
which serves data from a PostgreSQL database. 

Postman Docs [here](api.postman_collection.json)

## Overview

### Notable features

1. Two tables are used here: the `audio_shorts` table and `creators` table. `audio_shorts` has a foreign key `creator_id`
that refers to the `id` primary key field of `creators`. 
2. `status` field for both `audio_shorts` and `creators`, to provide useful metadata for internal usage. Other metadata 
   includes `created_at`, `updated_at`, and auto-incremented `id`.
3. `Delete` vs `HardDelete`: `Delete` changes `status` to 'deleted', whereas hard delete removes the entry from `audio_shorts`.
4. Pagination of audio shorts retrieval, by specifying `page` and `limit` in queries - page must be > 0 i.e. starts from 1.
5. Unit tests in Go, integration tests using Postman.
6. Libraries used: `gqlgen` for GraphQL, `golang-migrate` for migrations. No ORM for SQL, although did consider to use `GORM` or 
`SQLBoiler`
7. `Update` is a simple update, means all values must be specified in the mutation (no partial updates).

### Local Deployment

The service uses Docker to host the Go backend and the PostgreSQL database. The config can be found in the .env file.
[WIP] using some script to migrate and start docker-compose
