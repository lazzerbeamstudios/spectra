env "docker" {
  src = "ent://ent/schema"
  url = "postgres://postgres:Ehnz2a3AyrKqnurbXjx9@postgres-service:5432/postgres?sslmode=disable"
  dev = "docker://postgres/16/dev?search_path=public"
  migration {
    dir = "file://ent/migrate/migrations"
  }
}

env "dev" {
  src = "ent://ent/schema"
  url = "postgres://postgres:Ehnz2a3AyrKqnurbXjx9@localhost:5432/postgres?sslmode=disable"
  dev = "docker://postgres/16/dev?search_path=public"
  migration {
    dir = "file://ent/migrate/migrations"
  }
}
