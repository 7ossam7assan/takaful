package repositories

import (
	"Test/app"
	"Test/app/models"
	"fmt"
)
// todo code to interface later
//type UrgentCasesRepositoryRepository struct {
//}

func FindAll() []*models.UrgentCase{
	sql := "SELECT * FROM urgent_cases"
	rows, err := app.DB.Query(sql)
	defer rows.Close()

	var urgentCases []*models.UrgentCase
	if err != nil {
		panic(err)
	}
	urgentCase := models.UrgentCase{}

	for rows.Next() {
		var (
			Id   uint64
			Title string
			Description  string
			CreatedAt string
			UpdatedAt string
		)
		err = rows.Scan(&Id,&Title,&Description,&CreatedAt,&UpdatedAt)
		if err != nil {
			panic(err)
		}
		urgentCases = append(urgentCases, &models.UrgentCase{
			Id:Id,
			Title:Title,
			Description:Description,
			CreatedAt:CreatedAt,
			UpdatedAt:UpdatedAt,
		})
		fmt.Println(&urgentCase)
	}
	return urgentCases
}
