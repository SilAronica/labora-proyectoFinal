package services

func InsertLogInDb(dni string) {
	dbConnection.Exec("insert into logs (dni, estado, fecha_creación) values ($1, $2', now());", dni, "EN PROCESO")
}
