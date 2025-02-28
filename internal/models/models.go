package models

func Models() []interface{} {
	return []interface{}{
		&User{},
		&Note{},
	}
}
