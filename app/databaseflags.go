package app

type Database struct {
	DatabaseURI    string `short:"u" long:"database-uri" env:"DATABASE_URI" description:"database URI to connect to" default:"file:ent?mode=memory&cache=shared&_fk=1"`
	DatabaseDriver string `short:"s" long:"database-driver" env:"DATABASE_DRIVER" description:"Database driver to use" default:"sqlite3"`
}
