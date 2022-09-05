# Graphile Quickstart

Bootstrap a postgraphile based database project

- docker compose file for postgres 14
- initializes database structure for graphile migrate
- starts graphile-migrate and postgraphile

## Setup

```sh
source .env
make install
make up
make create
make init
make start
```

You now have postgraphile and graphile-migrate running.

## Create a database table

Open `migrations/current.sql` and add the following:

```sql
drop table if exists people;
create table people (
       id serial primary key,
       name text not null
);
```

## Test the generated GraphQL API

Visit http://localhost:5000/graphiql

Post to the generated createPerson mutation:

```gql
mutation {
  createPerson(input: { person: { name: "Ryan" } }) {
    person {
      id
      name
    }
  }
}
```

You should see in the result pane:

```json
{
  "data": {
    "createPerson": {
      "person": {
        "id": 4,
        "name": "Ryan"
      }
    }
  }
}
```

Post a query:

```gql
query {
  people {
    edges {
      node {
        id
        name
      }
    }
  }
}
```

See the result:

```gql
{
  "data": {
    "people": {
      "edges": [
        {
          "node": {
            "id": 1
            "name": "Ryan"
          }
        }
      ]
    }
  }
}
```

## View data in postgres

Connect to the database with psql:

```
$ make psql

app_development=> select * from people;
 id | name
----+-------
  1 | Ryan 
```

## Commit the migration


```
$ make commit
npx graphile-migrate commit
graphile-migrate[shadow]: dropped database 'app_development_shadow'
graphile-migrate[shadow]: recreated database 'app_development_shadow'
graphile-migrate[shadow]: Up to date — no committed migrations to run
graphile-migrate: New migration '000001.sql' created
graphile-migrate[shadow]: Running migration '000001.sql'
graphile-migrate[shadow]: 1 committed migrations executed
graphile-migrate: Running migration '000001.sql'
graphile-migrate: 1 committed migrations executed
```

The file current.sql will copied to `migrations/000001.sql`.  The next
migration can be now developed in `migrations/current.sql`.
