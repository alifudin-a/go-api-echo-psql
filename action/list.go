package action

import (
	"log"
	"net/http"

	"github.com/alifudin-a/go-api-echo-psql/database"
	"github.com/alifudin-a/go-api-echo-psql/model"
	"github.com/labstack/echo"
)

// GetEmployees : fetch all data from employee
func GetEmployees(c echo.Context) (err error) {

	db := database.OpenDB()

	sqlStatement := "SELECT * FROM employees ORDER BY id"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusCreated, rows)
	}

	defer rows.Close()

	emp := model.Employees{}

	for rows.Next() {
		employee := model.Employee{}
		err2 := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age)
		if err2 != nil {
			log.Println(err)
		}

		emp.Employees = append(emp.Employees, employee)
	}

	return c.JSON(http.StatusOK, emp)
}
