package services

// GetQuote gets a quote from the database
func GetQuote(id int) (string, error) {
	var quote string

	sql := "SELECT text FROM quote WHERE id = ?;"
	err := dbm.QueryRow(sql, id).Scan(&quote)
	if err != nil {
		return "", err
	}

	return quote, nil
}

// GetRandomQuote gets a random quote from the database
func GetRandomQuote() (string, error) {
	var quote string

	sql := "SELECT text FROM quote ORDER BY RANDOM() LIMIT 1;"
	err := dbm.QueryRow(sql).Scan(&quote)
	if err != nil {
		return "", err
	}

	return quote, nil
}

// GetQuotes gets a list of all the quotes from the database
func GetQuotes() ([]string, error) {
	var quotes []string
	var quote string

	sql := "SELECT text FROM quote ORDER BY id;"
	rows, err := dbm.QueryRows(sql)
	if err != nil {
		return quotes, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&quote)
		quotes = append(quotes, quote)
	}

	return quotes, nil
}
