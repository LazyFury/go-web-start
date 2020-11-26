package api

import (
	"github.com/Treblex/go-echo-demo/server/middleware"
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

var modelAppointment model.Appointment
var modelAppointmentLog model.AppointmentLog

func appointment(g *gin.RouterGroup) {
	modelAppointment.SetController(&modelAppointment)
	modelAppointmentGroup := g.Group("/appointment", middleware.Auth)
	modelAppointment.Install(modelAppointmentGroup, "")

	modelAppointmentLog.BaseControll.Model = &modelAppointmentLog
	modelAppointmentLogGroup := g.Group("/appointment-log", middleware.Auth)
	modelAppointmentLog.Install(modelAppointmentLogGroup, "")
}
