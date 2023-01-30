package server

import (
	"fxapp/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Handler for http requests
type Server struct {
	App    *fiber.App
	Logger *zap.SugaredLogger
	DB     *db.GormDatabase
}

// New http handler
func New(logger *zap.SugaredLogger, db *db.GormDatabase) *Server {
	app := fiber.New()
	app.Use(cors.New())

	server := Server{app, logger, db}
	return &server
}

/*
// RegisterRoutes for all http endpoints
func (h *Handler) registerRoutes() {
	h.mux.HandleFunc("/", h.hello)

}

func (h *Handler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	if r.URL.Query().Has("text") {
		h.add(w, r)
	} else if r.URL.Query().Has("id") {
		h.get(w, r)
	} else {
		w.Write([]byte("Hello World"))
	}

	h.logger.Info("Response done successfully")
}
func (h *Handler) add(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("add text")
	text := r.URL.Query().Get("text")
	id, _ := h.db.StoreText(text)
	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(int(id))))

}
func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	res, _ := h.db.GetTextByID(id)
	w.WriteHeader(200)
	w.Write([]byte(res))

}
*/
// Module provided to fx
var Module = fx.Options(
	fx.Provide(New),
)
