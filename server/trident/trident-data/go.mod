module trident-data

go 1.21

require (
	nvm-gocore v0.2.3
    nvm-sqlxe v0.2.3
	github.com/go-sql-driver/mysql v1.7.1 
	github.com/jmoiron/sqlx v1.3.5 //indirect
	github.com/lib/pq v1.10.9 
	github.com/mattn/go-sqlite3 v1.14.22 
)


replace nvm-gocore => ../../nvm-gocore
replace nvm-sqlxe => ../../nvm-sqlxe
