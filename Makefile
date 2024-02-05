migrations:
	goose postgres "host=localhost user=postgresDev database=fast_notes password=postgresDev sslmode=disable" up
migrationsStatus:
	goose postgres "host=localhost user=postgresDev database=fast_notes password=postgresDev sslmode=disable" status