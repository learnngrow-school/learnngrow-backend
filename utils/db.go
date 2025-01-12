package utils

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/db/repository"
)

func GetRepo(c *gin.Context) *repository.Queries {
	repoget, _ := c.Get("repo")
	repo, _ := repoget.(*repository.Queries)
	return repo
}
