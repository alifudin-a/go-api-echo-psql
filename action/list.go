package action

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alifudin-a/go-api-echo-psql/database"
	"github.com/alifudin-a/go-api-echo-psql/model"
	"github.com/labstack/echo"
)

// GetEmployees : fetch all data from employee
func GetEmployees(c echo.Context) (err error) {

	db := database.OpenDB()
	defer db.Close()

	limit, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if limit == 0 {
		limit = 10
	}

	sqlStatement := "SELECT * FROM employees ORDER BY id LIMIT $1"
	rows, err := db.Query(sqlStatement, limit)
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
