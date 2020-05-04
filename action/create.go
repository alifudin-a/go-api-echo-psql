package action

import (
	"net/http"

	"github.com/alifudin-a/go-api-echo-psql/database"
	"github.com/alifudin-a/go-api-echo-psql/model"
	"github.com/labstack/echo"
)

// CreateEmployee : action to create employee
func CreateEmployee(c echo.Context) (err error) {

	db := database.OpenDB()

	emp := new(model.Employee)
	if err = c.Bind(emp); err != nil {
		return err
	}

	sqlStatement := "INSERT INTO employees (id, name, salary, age) VALUES ($1, $2, $3, $4)"

	defer db.Close()

	_, err = db.Query(sqlStatement, emp.ID, emp.Name, emp.Salary, emp.Age)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, emp)
}
