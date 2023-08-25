package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Empleado struct {
	Id     int
	Nombre string
	Activo bool
}

var employees = sampleEmployees()

func main() {
	router := gin.Default()

	router.GET("/", HomePage)
	router.GET("/employees", AllEmployees)
	router.GET("/employees/:id", FindEmployee)
	router.GET("/employeesparam/:name/:status", AddEmployee)
	router.GET("/employeesactive", ActiveEmployees)
	router.Run(":8080")
}

func sampleEmployees() []Empleado {
	return []Empleado{{
		Id:     1,
		Nombre: "emp1",
		Activo: false,
	},
		{
			Id:     2,
			Nombre: "emp2",
			Activo: true,
		},
		{
			Id:     3,
			Nombre: "emp3",
			Activo: true,
		}}
}

func HomePage(ctxt *gin.Context) {
	ctxt.String(200, "¡Bienvenido a la Empresa Gophers!")
}

func FindEmployee(ctxt *gin.Context) {
	for _, employee := range employees {
		empId, _ := strconv.Atoi(ctxt.Param("id"))
		if employee.Id == empId {
			ctxt.String(200, "Información del empleado %s, nombre: %s, activo: %t", ctxt.Param("id"), employee.Nombre, employee.Activo)
			return
		}
	}
	ctxt.String(404, "Información del empleado ¡No existe!")
}

func AllEmployees(ctxt *gin.Context) {
	ctxt.JSON(http.StatusOK, employees)
}

func AddEmployee(ctxt *gin.Context) {
	status, _ := strconv.ParseBool(ctxt.Param("status"))
	newEmp := Empleado{
		Id:     len(employees) + 1,
		Nombre: ctxt.Param("name"),
		Activo: status,
	}
	employees = append(employees, newEmp)
	ctxt.JSON(http.StatusOK, employees[len(employees)-1])
}

func ActiveEmployees(ctxt *gin.Context) {
	var activeEmps []Empleado
	for _, employee := range employees {
		if employee.Activo {
			activeEmps = append(activeEmps, employee)
		}
	}
	ctxt.JSON(http.StatusOK, activeEmps)
}
