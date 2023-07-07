# commande pour creer la dB
` docker compose exec -it postgres createdb --username=auth_user --owner=auth_user auth_db `
# commande pour drop la db

` docker compose exec -it postgres dropdb auth_db`
# Pour run docker 

` Docker compose up `
# Pour run docker en détacher
` Docker compose up -d `

# Pour les migrations golang 

` github.com/golang-migrate/migrate `

# Pour créer les migrations 

` migrate create -ext sql -dir db/migrations make_user_table `

# Pour faire des migrations 

`migrate -path db/migrations -database "postgres://auth_user:rootme@localhost:5432/auth_db?sslmode=disable" --verbose up `