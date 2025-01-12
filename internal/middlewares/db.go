package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/utils"
)

var (
	pool *pgxpool.Pool
	once sync.Once
)

func getDbPool() *pgxpool.Pool {
	once.Do(func() {
		var err error
		pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatalf("Connection pool couldn't be created")
		}
	})
	return pool
}

func DbMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := getDbPool().Acquire(c)
		if err != nil {
			utils.Throw(c, http.StatusInternalServerError, err)
			return
		}
		defer conn.Release()

		repo := repository.New(conn)

		c.Set("repo", repo)
		c.Next()
	}
}
