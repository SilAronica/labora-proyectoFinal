package services

func InsertLogInDb(dni string) {
	dbConnection.Exec("insert into logs (dni, estado, fecha_creacion) values ($1, $2', now());", dni, "EN PROCESO")
}

func InsertWalletInDb(dni string) {
	dbConnection.Exec("insert into logs (dni, estado, fecha_creacion) values ($1, $2', now());", dni, "EN PROCESO")
}
