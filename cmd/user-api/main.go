package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-mongo-lab/config"
	"go-mongo-lab/handlers"
	"go-mongo-lab/routes"
)

func main() {
	// 1) Carrega env (PORT, MONGO_URI) com defaults
	cfg := config.Load()

	// 2) Conecta no Mongo usando a URI do env
	collection := config.ConnectDB(cfg.MongoURI)

	// 3) Injeta dependências (handler) e monta rotas
	h := handlers.NewUserHandler(collection)
	r := routes.SetupRouter(h) // /healthz e /v1/...

	// 4) Sobe o servidor HTTP
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("user-api ouvindo em http://localhost:%s\n", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("erro ao subir servidor: %v", err)
		}
	}()

	// 5) Graceful shutdown (SIGINT/SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("erro no shutdown: %v", err)
	}

	log.Println("servidor finalizado com sucesso")
}
