package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/labstack/echo/v4"
)

var modelAppointment model.Appointment
var modelAppointmentLog model.AppointmentLog

func appointment(g *echo.Group) {
	modelAppointment.BaseControll.Model = &modelAppointment
	modelAppointment.Install(g, "/appointment")

	modelAppointmentLog.BaseControll.Model = &modelAppointment
	modelAppointmentLog.Install(g, "/appointment-log")
}
