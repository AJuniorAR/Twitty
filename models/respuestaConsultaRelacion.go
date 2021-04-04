package models

//RespuestaConsultaRelacion tiene el true o falsw que se obtiene de consultar la relacion entre 2 usuarios
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
