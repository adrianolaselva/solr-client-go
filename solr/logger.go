package solr


type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}