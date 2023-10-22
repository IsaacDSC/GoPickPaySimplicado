package web

import (
	http_controllers "github.com/IsaacDSC/GoPickPaySimplicado/web/http/controllers"
	"github.com/go-chi/chi"
)

func StartRoutes(router *chi.Mux) {
	healthCheck := http_controllers.NewHealthCheckController("/")
	router.Get(healthCheck.Path, healthCheck.Index)

	router.Route("/api/v1", func(apiRouter chi.Router) {
		transactions := http_controllers.NewTransactionController(
			"/transaction",
		)
		apiRouter.Post(transactions.Path, transactions.PostTransaction)
	})

}
