package middleware

import (
	"github.com/Vinicamilotti/charlie/secrets"
	"github.com/gin-gonic/gin"
)

func AuthOwner(g *gin.Context) {
	getAuth := g.Request.Header.Get("Authorization")
	hash := secrets.GetSecrets().AuthHash
	if hash == getAuth {
		g.Next()
		return
	}

	g.AbortWithStatusJSON(401, gin.H{
		"error": "unauthorized",
	})
}
