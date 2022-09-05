postgraphile: cd database && npx postgraphile --watch --owner-connection $WATCH_DATABASE_URL --enhance-graphiql --append-plugins @graphile-contrib/pg-simplify-inflector --schema app_public
backend: cd backend && npx nodemon main.js
migrate: cd database && npx graphile-migrate watch
frontend: cd frontend-sk && npm run dev
