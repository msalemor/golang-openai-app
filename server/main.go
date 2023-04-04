package main

import (
	"os"
	"server/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
	dotenv "github.com/subosito/gotenv"
)

var app common.App

func LoadEnv() {
	dotenv.Load()

	stemp := os.Getenv("OPENAI_TEMPERATURE")
	temp, _ := strconv.ParseFloat(stemp, 64)

	sN := os.Getenv("OPENAI_N")
	N, _ := strconv.Atoi(sN)

	sMaxTokens := os.Getenv("OPENAI_MAX_TOKENS")
	maxTokens, _ := strconv.Atoi(sMaxTokens)

	app = common.App{
		APIKey:      os.Getenv("OPENAI_API_KEY"),
		Endpoint:    os.Getenv("OPENAI_ENDPOINT"),
		Temperature: temp,
		N:           N,
		MaxTokens:   maxTokens,
	}
	log.Info("Loaded environment variables", app)
}

func main() {
	LoadEnv()

	fiberApp := fiber.New()
	fiberApp.Use(logger.New())

	// Post for Azure OpenAI Post
	fiberApp.Post("/openai", func(c *fiber.Ctx) error {
		var request common.AppRequest
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		completion, err := common.AzureOpenAIRequest(request.Prompt, app)
		if err != nil {
			return err
		}
		log.Info(completion)
		response := common.AppResponse{
			Prompt:     request.Prompt,
			Completion: completion,
		}

		return c.JSON(response)
	})

	port := os.Getenv("APP_PORT")

	fiberApp.Static("/", "./public", fiber.Static{})
	log.Info("Starting server on port " + port)
	fiberApp.Listen(":3010")
}
