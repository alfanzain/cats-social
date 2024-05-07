## ERD

```
Table users {
  id int [pk, increment]
  email varchar [not null, unique]
  name varchar
  password varchar
  created_at datetime
  updated_at datetime
}

Table cats {
  id int [pk, increment]
  user_id int [ref: > users.id]
  name varchar(30) [not null]
  race enum(
    'PERSIAN',
    'MAINE_COON',
    'SIAMESE',
    'RAGDOLL',
    'BENGAL',
    'SPHYNX',
    'BRITISH_SHORTHAIR',
    'ABYSSINIAN',
    'SCOTTISH_FOLD',
    'BIRMAN'
  ) [not null]
  sex enum(
    'MALE', 
    'FEMALE'
  ) [not null]
  age_in_month int(6) [not null, default: 1]
  description varchar(200) [not null]
  has_matched boolean
  created_at datetime
  updated_at datetime
}

Table cat_images {
  id int [pk, increment]
  cat_id int [not null, ref: > cats.id]
  image_url varchar [not null]
  created_at datetime
  updated_at datetime
}

Table cat_matches {
  issuer_cat_id int [not null, ref: > cats.id]
	match_cat_id int [not null, ref: > cats.id]
  status enum(
    'PENDING',
    'APPROVED'
  ) [not null, default: 'PENDING']
  created_at datetime
  updated_at datetime
}
```

### What We Did

#### Docker

First time

- `docker compose up --build` for idk seriously

#### Migrations
- migrate create -ext sql -dir db/migrations -seq create_users_table
- migrate -database postgres://postgres:password@postgres:5432/local_cats_social?sslmode=disable -path db/migrations up

When migration failed
- `psql -d local_cats_social -U postgres`
- `DROP TABLE users;`
- `TRUNCATE schema_migrations;`