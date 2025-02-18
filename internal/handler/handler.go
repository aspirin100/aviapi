package handler

type Handler struct {
	AirflightManager
}

func New(manager AirflightManager) *Handler {
	return &Handler{
		AirflightManager: manager,
	}
}
