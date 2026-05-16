package middleware

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"time"
	"vaultex/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthMiddleware(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		signature := c.GetHeader("x-signature")
		timestamp := c.GetHeader("x-timestamp")

		// 1. Check all headers present
		if apiKey == "" || signature == "" || timestamp == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing required headers",
			})
			return
		}

		// 2. Check timestamp — reject if older than 5 minutes
		ts, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil || time.Now().Unix()-ts > 300 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "request expired",
			})
			return
		}

		// 3. Lookup org by api_key
		var encryptedSecret string
		var orgID string
		var isActive bool
		err = db.QueryRow(c, `
            SELECT id, api_secret, is_active 
            FROM organizations 
            WHERE api_key = $1
        `, apiKey).Scan(&orgID, &encryptedSecret, &isActive)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid api key",
			})
			return
		}

		// 4. Check org is active
		if !isActive {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "organization is inactive",
			})
			return
		}

		// 5. Decrypt api_secret
		secret, err := util.DecryptSecret(encryptedSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "server error",
			})
			return
		}

		// 6. Read body and restore it for next handler
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 7. Verify HMAC
		valid := util.VerifyHMAC(
			timestamp,
			c.Request.Method,
			c.Request.URL.Path,
			string(bodyBytes),
			secret,
			signature,
		)
		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid signature",
			})
			return
		}

		// 8. Attach org id to context for handlers
		c.Set("org_id", orgID)
		c.Next()
	}
}
