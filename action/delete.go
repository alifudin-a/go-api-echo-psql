package action

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/alifudin-a/go-api-echo-psql/database"
	"github.com/labstack/echo"
)

// DeleteEmployee : action for delete employee
func DeleteEmployee(c echo.Context) (err error) {

	db := database.OpenDB()
	defer db.Close()

	id := c.Param("id")

	sqlStatement := "DELETE FROM employees WHERE id=$1"
	_, err = db.Query(sqlStatement, id)
	if err != nil {
		fmt.Println("No rows were returned!", sql.ErrNoRows)
	}

	return c.JSON(http.StatusOK, "Deleted")
}
