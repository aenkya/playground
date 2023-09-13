package db

type DB interface{}

type Database struct {
	db DB
}

func Practice() DB {
	return MongoDB()
}

func MongoDB() *Database {
	return &Database{
		db: initMongoDB(),
	}
}

func initMongoDB() DB {
	return nil
}
