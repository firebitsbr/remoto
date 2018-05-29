// Code generated by Remoto; DO NOT EDIT.

// Package example contains the HTTP server for example services.
package example

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/machinebox/remoto/go/remotohttp"
	"github.com/pkg/errors"
	
	"github.com/machinebox/remoto/remototypes"	
	
)

// Run is the simplest way to run the services.
func Run(addr string,
	facebox Facebox,
) error {
	server := New(
		facebox,
	)
	if err := server.Describe(os.Stdout); err != nil {
		return errors.Wrap(err, "describe service")
	}
	if err := http.ListenAndServe(addr, server); err != nil {
		return err
	}
	return nil
}

// New makes a new remotohttp.Server with the specified services
// registered.
func New(
	facebox Facebox,
) *remotohttp.Server {
	server := &remotohttp.Server{
		OnErr: func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Fprintf(os.Stderr, "%s %s: %s\n", r.Method, r.URL.Path, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}
	
	RegisterFaceboxServer(server, facebox)
	return server
}

// 
type TeachResponse struct {
	// Error is an error message if one occurred.
	Error string `json:"error"`
	
}

// 
type CheckRequest struct {
	// 
	Image remototypes.File `json:"image"`
	
}

// 
type Faces struct {
	// 
	Name string `json:"name"`
	// 
	Matched bool `json:"matched"`
	
}

// 
type CheckResponse struct {
	// 
	Faces Faces `json:"faces"`
	// Error is an error message if one occurred.
	Error string `json:"error"`
	
}

// 
type TeachFile struct {
	// 
	Image remototypes.File `json:"image"`
	
}

// 
type TeachRequest struct {
	// 
	Name string `json:"name"`
	// 
	TeachFiles TeachFile `json:"teach_files"`
	
}



// Facebox provides facial detection and recognition capabilities.
type Facebox interface {
	// 
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
// 
	Teach(context.Context, *TeachRequest) (*TeachResponse, error)

}

// RegisterFaceboxServer registers a Facebox with a remotohttp.Server.
func RegisterFaceboxServer(server *remotohttp.Server, service Facebox) {
	srv := &httpFaceboxServer{
		service: service,
		server: server,
	}
	
	server.Register("/remoto/Facebox.Check", http.HandlerFunc(srv.Check))
	
	server.Register("/remoto/Facebox.Teach", http.HandlerFunc(srv.Teach))
	
}

type httpFaceboxServer struct {
	// service is the Facebox being exposed by this
	// server.
	service Facebox
	// server is the remotohttp.Server that this server is
	// registered with.
	server *remotohttp.Server
}


// Check is an http.Handler wrapper for Facebox.Check.
func (srv *httpFaceboxServer) Check(w http.ResponseWriter, r *http.Request) {
	var reqs []*CheckRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}
	resps := make([]CheckResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.Check(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}
}
// Teach is an http.Handler wrapper for Facebox.Teach.
func (srv *httpFaceboxServer) Teach(w http.ResponseWriter, r *http.Request) {
	var reqs []*TeachRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}
	resps := make([]TeachResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.Teach(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}
} 
