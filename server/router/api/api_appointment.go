package api

import (
	"github.com/Treblex/go-echo-demo/server/middleware"
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/labstack/echo/v4"
)

var modelAppointment model.Appointment
var modelAppointmentLog model.AppointmentLog

func appointment(g *echo.Group) {
	modelAppointment.BaseControll.Model = &modelAppointment
	modelAppointmentGroup := g.Group("/appointment", middleware.UserJWT)
	modelAppointment.Install(modelAppointmentGroup, "")

	modelAppointmentLog.BaseControll.Model = &modelAppointmentLog
	modelAppointmentLogGroup := g.Group("/appointment-log", middleware.UserJWT)
	modelAppointmentLog.Install(modelAppointmentLogGroup, "")
}
