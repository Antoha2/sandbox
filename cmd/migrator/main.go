package main

func main() {
	//var storagePath, migrationsPath, migrationsTable string

	// flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	// flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	// flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	// flag.Parse()

	// if storagePath == "" {
	// 	panic("storage-path is required")
	// }
	// if migrationsPath == "" {
	// 	panic("migrations-path is required")
	// }

	// connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
	// 	cfg.DBConfig.User,
	// 	cfg.DBConfig.Password,
	// 	cfg.DBConfig.Host,
	// 	cfg.DBConfig.Port,
	// 	cfg.DBConfig.Dbname,
	// 	cfg.DBConfig.Sslmode,
	// )

	//dbType, dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	// m, err := migrate.New(
	// 	"file://"+migrationsPath,
	// 	fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	// )
	// m, err := migrate.New("file://"+migrationsPath,
	// fmt.Sprintf("%v://%v:%v/%v&sslmode=%v",
	// %s?x-migrations-table=%s", storagePath, migrationsTable),
	// if err != nil {
	// 	panic(err)
	// }

	// if err := m.Up(); err != nil {
	// 	if errors.Is(err, migrate.ErrNoChange) {
	// 		fmt.Println("no migrations to apply")

	// 		return
	// 	}

	// 	panic(err)
	// }

	// fmt.Println("migrations applied")
}
