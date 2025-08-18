package handler

import (
	"net/http"

	"github.com/006lp/akashchat-api-go/internal/model"
	"github.com/006lp/akashchat-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

// ChatHandler handles chat-related HTTP requests
type ChatHandler struct {
	sessionService *service.SessionService
	akashService   *service.AkashService
}

// NewChatHandler creates a new ChatHandler instance
func NewChatHandler(sessionService *service.SessionService, akashService *service.AkashService) *ChatHandler {
	return &ChatHandler{
		sessionService: sessionService,
		akashService:   akashService,
	}
}

// ChatCompletions handles the /v1/chat/completions endpoint
func (h *ChatHandler) ChatCompletions(c *gin.Context) {
	var req model.ChatCompletionRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code: 400,
			Data: model.ErrorData{Message: "Invalid request format: " + err.Error()},
		})
		return
	}

	// Get session token
	sessionToken, err := h.sessionService.GetSessionToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Code: 500,
			Data: model.ErrorData{Message: "Failed to get session token: " + err.Error()},
		})
		return
	}

	// Process chat request
	if req.Model == "AkashGen" {
		// Set default values
		temperature := 0.85
		if req.Temperature != nil {
			temperature = *req.Temperature
		}

		topP := 1.0
		if req.TopP != nil {
			topP = *req.TopP
		}

		// Handle image generation
		data, err := h.akashService.ProcessImageGeneration(req, sessionToken, temperature, topP)
		if err != nil {
			if err.Error() == "invalid model" {
				c.JSON(http.StatusInternalServerError, model.APIResponse{
					Code: 500,
					Data: model.ErrorData{Message: "Error Model."},
				})
				return
			}
			c.JSON(http.StatusInternalServerError, model.APIResponse{
				Code: 500,
				Data: model.ErrorData{Message: "Image generation failed: " + err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, model.APIResponse{
			Code: 200,
			Data: data,
		})
	} else {
		// Set default values
		temperature := 0.6
		if req.Temperature != nil {
			temperature = *req.Temperature
		}

		topP := 0.95
		if req.TopP != nil {
			topP = *req.TopP
		}

		// Handle text generation
		data, err := h.akashService.ProcessTextGeneration(req, sessionToken, temperature, topP)
		if err != nil {
			if err.Error() == "invalid model" {
				c.JSON(http.StatusInternalServerError, model.APIResponse{
					Code: 500,
					Data: model.ErrorData{Message: "Error Model."},
				})
				return
			}
			c.JSON(http.StatusInternalServerError, model.APIResponse{
				Code: 500,
				Data: model.ErrorData{Message: "Text generation failed: " + err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
