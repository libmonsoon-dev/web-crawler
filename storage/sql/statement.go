package sql

type Statement interface {
	// Sql returns parametrized sql query with list of arguments.
	Sql() (query string, args []interface{})
	// DebugSql returns debug query where every parametrized placeholder is replaced with its argument.
	// Do not use it in production. Use it only for debug purposes.
	DebugSql() (query string)
}
