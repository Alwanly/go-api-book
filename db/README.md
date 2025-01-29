# Database Migrations

This document provides steps to run database migrations using Atlas.

## Prerequisites

1. Ensure you have Go installed on your machine.
2. Install Atlas by running the following command:
    ```sh
    curl -sSf https://atlasgo.sh | sh
    ```

## Setup

1. Clone the repository and navigate to the `db` directory:
    ```sh
    git clone <repository-url>
    cd <repository-directory>/db
    ```

2. Install Go dependencies:
    ```sh
    make install
    ```

3. Ensure the `.env.db` file is correctly configured with your database URI:
    ```plaintext
    DB_URI=postgres://local_test:local_test@localhost:5432/db_local_test?sslmode=disable
    ```

## Running Migrations

### Create a New Migration

To create a new migration file, run:
```sh
make new NAME=<migration_name>
```

### Apply Migrations

To apply all pending migrations, run:
```sh
make up
```

### Check Migration Status

To check the current status of migrations, run:
```sh
make status
```

### Generate Migration Diff

To generate a diff for the current database schema, run:
```sh
make diff
```

### Reset the Database

To reset the database and reapply all migrations, run:
```sh
make reset
```

### Hash Migrations

To hash the migrations, run:
```sh
make hash
```

## Writing Schema

### Create a New Schema File

1. Navigate to the `schema` directory:
    ```sh
    cd schema
    ```

2. Create a new schema file with the `.hcl` extension, for example:
    ```sh
    touch new_schema.hcl
    ```

3. Define your schema in the new file. For example:
    ```hcl
    table "users" {
      schema = schema.public
      column "id" {
        type = int
      }
      column "username" {
        type = varchar(36)
      }
      column "password" {
        type = varchar(128)
      }
      primary_key {
        columns = [column.id]
      }
    }
    ```

### Add Schema Config to `atlas.hcl`

1. Open the `atlas.hcl` file located in the `db` directory.
2. Add the new schema file to the `src` array under the `env "local"` section. For example:
    ```hcl
    env "local" {
      url = var.url
      dev = "docker://postgres/14-alpine/kucingmenangis?search_path=public&sslmode=disable"
      src = [
        "file://schema/codebase.hcl",
        "file://schema/new_schema.hcl",  // Add your new schema file here
      ]

      migration {
        dir = "file://migration"
        revisions_schema = var.revisions_schema
        format = atlas
      }
    }
    ```

3. Save the changes to `atlas.hcl`.

## Additional Information

For more details on Atlas, visit the [Atlas documentation](https://atlasgo.io/).
