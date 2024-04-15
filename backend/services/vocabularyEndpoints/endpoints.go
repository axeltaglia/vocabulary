package vocabularyEndpoints

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"vocabulary/entities"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
)

type Endpoints struct {
	router              *gin.Engine
	txRepositoryHandler entities.TxRepositoryHandler
}

func (o *Endpoints) handle() {
	o.handleTxWithVocabularyEntity("/getVocabularies", o.getVocabularies)
	o.handleTxWithVocabularyEntity("/getVocabulary/:id", o.getVocabulary)
	o.handleTxWithVocabularyEntity("/getVocabularyCategories/:id", o.getVocabularyCategories)
	o.handleTxWithVocabularyEntity("/updateVocabularyWithCategories", o.updateVocabularyWithCategories)
	o.handleTxWithVocabularyEntity("/deleteVocabulary/:id", o.deleteVocabulary)
	o.handleTxWithVocabularyEntity("/createVocabularyWithCategories", o.createVocabularyWithCategories)
	o.handleTxWithVocabularyEntity("/getCategories", o.getCategories)
}

func (o *Endpoints) handleTxWithVocabularyEntity(relativePath string, f func(c *gin.Context, vocabularyEntity VocabularyEntity.Entity)) {
	o.router.POST(relativePath, func(c *gin.Context) {
		txRepositoryFactory := o.txRepositoryHandler.GetTxRepositoryFactory()

		defer func() {
			if r := recover(); r != nil {
				txRepositoryFactory.RollbackTransaction()
				stackTrace := string(debug.Stack())
				fmt.Printf("%v\n%s\n", r, stackTrace)
			}
		}()

		vocabularyRepository := txRepositoryFactory.GetVocabularyRepository()
		vocabularyEntity := VocabularyEntity.New(vocabularyRepository)
		f(c, vocabularyEntity)

		if txRepositoryFactory.TransactionError() != nil {
			txRepositoryFactory.RollbackTransaction()
		} else {
			txRepositoryFactory.CommitTransaction()
		}
	})
}

func (o *Endpoints) ListenAndServe(apiPort string) {
	o.handle()
	err := http.ListenAndServe(":"+apiPort, o.router)
	if err != nil {
		panic(err)
	}
}

func NewEndpoints(txRepositoryHandler entities.TxRepositoryHandler) *Endpoints {
	router := getGinRouter()
	return &Endpoints{
		router:              router,
		txRepositoryHandler: txRepositoryHandler,
	}
}

func getGinRouter() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.Use(loggerMiddleware())
	router.Use(authenticatedMiddleware())
	return router
}

func corsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	return cors.New(corsConfig)
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		data := map[string]interface{}{
			"status":    strconv.FormatInt(int64(statusCode), 10),
			"latency":   strconv.FormatInt(int64(latency), 10),
			"client_ip": clientIP,
			"method":    method,
			"path":      path,
			"raw_query": raw,
			"error":     errorMessage,
		}

		logger.GetLogger().LogWithFields(data)
	}
}

func authenticatedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method here
			return []byte("your-secret-key"), nil // Replace "your-secret-key" with your actual secret key
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// If the token is valid, continue with the next middleware
		c.Next()
	}
}
