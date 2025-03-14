package requests

// func getCategories() ([]Categories, error) {
// 	rows, err := helpers.DataBase.Query("SELECT id, name FROM Category")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var categories []Categories
// 	for rows.Next() {
// 		var c Categories
// 		if err := rows.Scan(&c.ID, &c.Name); err != nil {
// 			return nil, err
// 		}
// 		categories = append(categories, c)
// 	}
// 	return categories, rows.Err()
// }
