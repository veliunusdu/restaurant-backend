package middleware

import (
"net/http"
"strings"

"golang-restaurant-management/helpers"

"github.com/gin-gonic/gin"
)

// Authentication validates the Authorization token and sets user claims in context.
func Authentication() gin.HandlerFunc {
return func(c *gin.Context) {
authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
if authHeader == "" {
c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
c.Abort()
return
}

clientToken := authHeader
if parts := strings.SplitN(authHeader, " ", 2); len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
clientToken = strings.TrimSpace(parts[1])
}

if clientToken == "" {
c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
c.Abort()
return
}

claims, msg := helpers.ValidateToken(clientToken)
if msg != "" {
c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
c.Abort()
return
}

c.Set("email", claims.Email)
c.Set("first_name", claims.First_name)
c.Set("last_name", claims.Last_name)
c.Set("uid", claims.Uid)

c.Next()
}
}
