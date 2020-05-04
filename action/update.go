package action

import (
	"net/http"

	"github.com/alifudin-a/go-api-echo-psql/database"
	"github.com/alifudin-a/go-api-echo-psql/model"
	"github.com/labstack/echo"
)

// UpdateEmployee : action for update a employee
func UpdateEmployee(c echo.Context) (err error) {

	db := database.OpenDB()
	defer db.Close()

	emp := new(model.Employee)
	if err = c.Bind(emp); err != nil {
		return err
	}

	sqlStatement := "UPDATE employees SET name=$1, salary=$2, age=$3 WHERE id=$4"
	_, err = db.Query(sqlStatement, emp.Name, emp.Salary, emp.Age, emp.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, emp)
}
