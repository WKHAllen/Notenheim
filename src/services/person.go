package services

// GetPerson gets a person from the database
func GetPerson(id int) (string, error) {
	var firstname, lastname string

	sql := "SELECT firstname, lastname FROM person WHERE id = ?;"
	err := dbm.QueryRow(sql, id).Scan(&firstname, &lastname)
	if err != nil {
		return "", err
	}

	return firstname + " " + lastname, nil
}

// GetRandomPerson gets a random person from the database
func GetRandomPerson() (string, error) {
	var firstname, lastname string

	sql := "SELECT firstname, lastname FROM person ORDER BY RANDOM() LIMIT 1;"
	err := dbm.QueryRow(sql).Scan(&firstname, &lastname)
	if err != nil {
		return "", err
	}

	return firstname + " " + lastname, nil
}

// GetPeople gets a list of all the people from the database
func GetPeople() ([]string, error) {
	var people []string
	var firstname, lastname string

	sql := "SELECT firstname, lastname FROM person ORDER BY id;"
	rows, err := dbm.QueryRows(sql)
	if err != nil {
		return people, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&firstname, &lastname)
		people = append(people, firstname + " " + lastname)
	}

	return people, nil
}
