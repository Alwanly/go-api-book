schema "public" {}

table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = varchar(36)
  }
  column "username" {
    null = false
    type = varchar(36)
  }
  column "password" {
    null = false
    type = integer
  }
  column "created_at" {
    null = true
    type = bigint
  }
  column "created_by" {
    null = true
    type = varchar(36)
  }
  column "updated_at" {
    null = true
    type = bigint
  }
  column "updated_by" {
    null = true
    type = varchar(36)
  }
  primary_key {
    columns = [column.id]
  }
}