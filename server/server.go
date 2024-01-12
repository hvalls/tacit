package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tacit/config"
	"tacit/handler"
)

type Server struct {
	port string
}

func New(port string) *Server {
	return &Server{port}
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
	fmt.Println("Registering endpoint: ", e)

	http.HandleFunc(e.Path, func(w http.ResponseWriter, r *http.Request) {
		//TODO: check http method
		args, err := buildArgs(r, e.Args)
		if err != nil {
			panic(err) //TODO: change this
		}
		stdout, stderr, err := handler.Handle(handler.DEFAULT_SHELL, e.Handler, args)
		if err != nil {
			panic(err) //TODO: change this
		}

		w.Header().Add("Content-Type", "application/json")

		if stderr != "" {
			e := &ErrorResponse{stderr}
			jsonData, err := json.Marshal(e)
			if err != nil {
				panic(err) //TODO: change this
			}
			w.WriteHeader(500)
			fmt.Fprint(w, string(jsonData))
			return
		}

		fmt.Fprint(w, stdout)
	})

	return nil
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
	fmt.Println("Ready. Tacit server is listening on port", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
