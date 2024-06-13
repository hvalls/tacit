package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tacit/config"
	"tacit/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r *gin.Engine
}

func New() *Server {
	r := gin.Default()
	return &Server{r}
}

func (s *Server) RegisterEndpoints(ee []config.Endpoint) error {
	for _, e := range ee {
		err := s.registerEndpoint(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) registerEndpoint(e config.Endpoint) error {
	switch e.Method {
	case http.MethodGet:
		s.r.GET(e.Path, func(c *gin.Context) {
			handleGet(e, c)
		})
	case http.MethodPost:
		s.r.POST(e.Path, func(c *gin.Context) {
			handlePost(e, c)
		})
	}
	return nil
}

func handleGet(e config.Endpoint, c *gin.Context) {
	args, err := buildArgs(c.Request, e.Args)
	if err != nil {
		panic(err) //TODO: change this
	}
	fmt.Printf("Attempting to handle | Shell: %s | Scritp Path: %s | Args: %v\n", handler.DEFAULT_SHELL, e.Handler, args)
	stdout, stderr, err := handler.Handle(handler.DEFAULT_SHELL, e.Handler, args)
	if err != nil {
		panic(err) //TODO: change this
	}

	if stderr != "" {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{stderr})
		return
	}

	c.JSON(http.StatusOK, newResponse(stdout))
}

func handlePost(e config.Endpoint, c *gin.Context) {
	//TODO: Implement
	c.JSON(http.StatusOK, &ErrorResponse{"not supported yet"})
}

func buildArgs(r *http.Request, configArgs []string) ([]string, error) {
	var args []string
	for _, carg := range configArgs {
		if strings.Contains(carg, "$query.") {
			queryParamName := strings.Split(carg, "$query.")[1]
			queryParamValue := r.URL.Query()[queryParamName]
			args = append(args, queryParamValue...)
			continue
		}
		args = append(args, carg)
	}

	return args, nil
}

func (s *Server) Listen() error {
	return s.r.Run()
}

type Response struct {
	Data map[string]any `json:"data"`
}

func newResponse(rawData string) *Response {
	var data map[string]any
	err := json.Unmarshal([]byte(rawData), &data)
	if err != nil {
		panic(err) //TODO: Change this
	}
	return &Response{data}
}

type ErrorResponse struct {
	Error string `json:"error"`
}
