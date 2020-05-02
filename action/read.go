package action

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/alifudin-a/go-api-echo-psql/database"
	"github.com/alifudin-a/go-api-echo-psql/model"
	"github.com/labstack/echo"
)

// GetEmployee : action for read employee by id
func GetEmployee(c echo.Context) (err error) {

	db := database.OpenDB()

	id := c.Param("id")
	emp := model.Employee{}

	sqlStatement := "SELECT * FROM employees WHERE id=$1"

	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&emp.ID, &emp.Name, &emp.Salary, &emp.Age)
	if err != nil {
		fmt.Println("No rows were returned!", sql.ErrNoRows)
	}

	return c.JSON(http.StatusOK, emp)
}
