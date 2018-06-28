// Code generated by Remoto; DO NOT EDIT.

// Package machinebox contains the HTTP server for machinebox services.
package machinebox

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/matryer/remoto/go/remotohttp"
	"github.com/matryer/remoto/go/remotohttp/remototypes"
	"github.com/pkg/errors"
)

// Facebox provides facial detection and recognition in images.
type Facebox interface {

	// CheckFaceprint checks to see if a Faceprint matches any known
	// faces.
	CheckFaceprint(context.Context, *CheckFaceprintRequest) (*CheckFaceprintResponse, error)

	// CheckFile checks an image file for faces.
	CheckFile(context.Context, *CheckFileRequest) (*CheckFileResponse, error)

	// CheckURL checks a hosted image file for faces.
	CheckURL(context.Context, *CheckURLRequest) (*CheckURLResponse, error)

	// FaceprintCompare compares faceprints to a specified target describing
	// similarity.
	FaceprintCompare(context.Context, *FaceprintCompareRequest) (*FaceprintCompareResponse, error)

	// GetState gets the Facebox state file.
	GetState(context.Context, *GetStateRequest) (*remototypes.FileResponse, error)

	// PutState sets the Facebox state file.
	PutState(context.Context, *PutStateRequest) (*PutStateResponse, error)

	// RemoveID removes a face with the specified ID.
	RemoveID(context.Context, *RemoveIDRequest) (*RemoveIDResponse, error)

	// Rename changes a person&#39;s name.
	Rename(context.Context, *RenameRequest) (*RenameResponse, error)

	// RenameID changes the name of a previously taught face, by ID.
	RenameID(context.Context, *RenameIDRequest) (*RenameIDResponse, error)

	// SimilarFile checks for similar faces from the face in an image file.
	SimilarFile(context.Context, *SimilarFileRequest) (*SimilarFileResponse, error)

	// SimilarID checks for similar faces by ID.
	SimilarID(context.Context, *SimilarIDRequest) (*SimilarIDResponse, error)

	// SimilarURL checks for similar faces in a hosted image file.
	SimilarURL(context.Context, *SimilarURLRequest) (*SimilarURLResponse, error)

	// TeachFaceprint teaches Facebox about a face from a Faceprint.
	TeachFaceprint(context.Context, *TeachFaceprintRequest) (*TeachFaceprintResponse, error)

	// TeachFile teaches Facebox a new face from an image file.
	TeachFile(context.Context, *TeachFileRequest) (*TeachFileResponse, error)

	// TeachURL teaches Facebox a new face from an image on the web.
	TeachURL(context.Context, *TeachURLRequest) (*TeachURLResponse, error)
}

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
		NotFound: http.NotFoundHandler(),
	}

	RegisterFaceboxServer(server, facebox)
	return server
}

// RegisterFaceboxServer registers a Facebox with a remotohttp.Server.
func RegisterFaceboxServer(server *remotohttp.Server, service Facebox) {
	srv := &httpFaceboxServer{
		service: service,
		server:  server,
	}
	server.Register("/remoto/Facebox.CheckFaceprint", http.HandlerFunc(srv.handleCheckFaceprint))
	server.Register("/remoto/Facebox.CheckFile", http.HandlerFunc(srv.handleCheckFile))
	server.Register("/remoto/Facebox.CheckURL", http.HandlerFunc(srv.handleCheckURL))
	server.Register("/remoto/Facebox.FaceprintCompare", http.HandlerFunc(srv.handleFaceprintCompare))
	server.Register("/remoto/Facebox.GetState", http.HandlerFunc(srv.handleGetState))
	server.Register("/remoto/Facebox.PutState", http.HandlerFunc(srv.handlePutState))
	server.Register("/remoto/Facebox.RemoveID", http.HandlerFunc(srv.handleRemoveID))
	server.Register("/remoto/Facebox.Rename", http.HandlerFunc(srv.handleRename))
	server.Register("/remoto/Facebox.RenameID", http.HandlerFunc(srv.handleRenameID))
	server.Register("/remoto/Facebox.SimilarFile", http.HandlerFunc(srv.handleSimilarFile))
	server.Register("/remoto/Facebox.SimilarID", http.HandlerFunc(srv.handleSimilarID))
	server.Register("/remoto/Facebox.SimilarURL", http.HandlerFunc(srv.handleSimilarURL))
	server.Register("/remoto/Facebox.TeachFaceprint", http.HandlerFunc(srv.handleTeachFaceprint))
	server.Register("/remoto/Facebox.TeachFile", http.HandlerFunc(srv.handleTeachFile))
	server.Register("/remoto/Facebox.TeachURL", http.HandlerFunc(srv.handleTeachURL))

}

// CheckFaceprintRequest is the request object for CheckFaceprint calls.
type CheckFaceprintRequest struct {

	// Faceprints is a list of Faceprints to check.
	Faceprints []string `json:"faceprints"`
}

// CheckFaceprintResponse is the response object for CheckFaceprint calls.
type CheckFaceprintResponse struct {

	// Faces is a list of faces checked from Faceprints.
	Faces []FaceprintFace `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// CheckFileRequest is the request object for CheckFile calls.
type CheckFileRequest struct {

	// File is the image to check for faces.
	File remototypes.File `json:"file"`
}

// CheckFileResponse is the response object for CheckFile calls.
type CheckFileResponse struct {

	// Faces is a list of faces that were found.
	Faces []Face `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// CheckURLRequest is the request object for CheckURL calls.
type CheckURLRequest struct {

	// URL is the address of the image to check.
	URL string `json:"url"`
}

// CheckURLResponse is the response object for CheckURL calls.
type CheckURLResponse struct {

	// Faces is a list of faces that were found.
	Faces []Face `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// Face describes a face.
type Face struct {

	// ID is the identifier of the source that was matched.
	ID string `json:"id"`

	// Name is the name of the identified person.
	Name string `json:"name"`

	// Matched is whether the face was recognized or not.
	Matched bool `json:"matched"`

	// Faceprint is the Facebox Faceprint of this face.
	Faceprint string `json:"faceprint"`

	// Rect is where the face appears in the source image.
	Rect Rect `json:"rect"`
}

// FaceprintCompareRequest is the request object for FaceprintCompare calls.
type FaceprintCompareRequest struct {

	// Target is the target Faceprint to which the Faceprints will be compared.
	Target string `json:"target"`

	// Faceprints is a list of Faceprints that will be compared to Target.
	Faceprints []string `json:"faceprints"`
}

// FaceprintCompareResponse is the response object for FaceprintCompare calls.
type FaceprintCompareResponse struct {

	// Confidences is a list of confidence values.
	// The order matches the order of FaceprintCompareRequest.Faceprints.
	Confidences []float64 `json:"confidences"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// FaceprintFace is a face.
type FaceprintFace struct {

	// Matched is whether the face was recognized or not.
	Matched bool `json:"matched"`

	// Confidence is a numerical value of how confident the AI
	// is that this is a match.
	Confidence float64 `json:"confidence"`

	// ID is the identifier of the source that matched.
	ID string `json:"id"`

	// Name is the name of the person recognized.
	Name string `json:"name"`
}

// GetStateRequest is the request object for GetState calls.
type GetStateRequest struct {
}

// PutStateRequest is the request object for PutState calls.
type PutStateRequest struct {

	// StateFile is the Facebox state file to set.
	StateFile remototypes.File `json:"state_file"`
}

// PutStateResponse is the response object for PutState calls.
type PutStateResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// Rect is a bounding box describing a rectangle of an image.
type Rect struct {

	// Top is the starting Y coordinate.
	Top int `json:"top"`

	// Left is the starting X coordinate.
	Left int `json:"left"`

	// Width is the width.
	Width int `json:"width"`

	// Height is the height.
	Height int `json:"height"`
}

// RemoveIDRequest is the request object for RemoveID calls.
type RemoveIDRequest struct {

	// ID is the identifier of the source to remove.
	ID string `json:"id"`
}

// RemoveIDResponse is the response object for RemoveID calls.
type RemoveIDResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// RenameIDRequest is the request object for RenameID calls.
type RenameIDRequest struct {

	// ID is the identifier of the source to rename.
	ID string `json:"id"`

	// Name is the new name to assign to the item matching ID.
	Name string `json:"name"`
}

// RenameIDResponse is the response object for RenameID calls.
type RenameIDResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// RenameRequest is the request object for Rename calls.
type RenameRequest struct {

	// From is the original name.
	From string `json:"from"`

	// To is the new name.
	To string `json:"to"`
}

// RenameResponse is the response object for Rename calls.
type RenameResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// SimilarFace is a detected face with similar matching faces.
type SimilarFace struct {

	// Rect is where the face appears in the image.
	Rect Rect `json:"rect"`

	// SimilarFaces is a list of similar faces.
	SimilarFaces []Face `json:"similar_faces"`
}

// SimilarFileRequest is the request object for SimilarFile calls.
type SimilarFileRequest struct {
	File remototypes.File `json:"file"`
}

// SimilarFileResponse is the response object for SimilarFile calls.
type SimilarFileResponse struct {
	Faces []SimilarFace `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// SimilarIDRequest is the request object for SimilarID calls.
type SimilarIDRequest struct {

	// ID is the identifier of the source to look for similar faces of.
	ID string `json:"id"`
}

// SimilarIDResponse is the response object for SimilarID calls.
type SimilarIDResponse struct {

	// Faces is a list of similar faces.
	Faces []SimilarFace `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// SimilarURLRequest is the request object for SimilarURL calls.
type SimilarURLRequest struct {
	URL string `json:"url"`
}

// SimilarURLResponse is the response object for SimilarURL calls.
type SimilarURLResponse struct {
	Faces []SimilarFace `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// TeachFaceprintRequest is the request object for TeachFaceprint calls.
type TeachFaceprintRequest struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Faceprint string `json:"faceprint"`
}

// TeachFaceprintResponse is the response object for TeachFaceprint calls.
type TeachFaceprintResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// TeachFileRequest is the request object for TeachFile calls.
type TeachFileRequest struct {

	// ID is an identifier describing the source, for example the filename.
	ID string `json:"id"`

	// Name is the name of the person in the image.
	Name string `json:"name"`

	// File is the image containing the face to teach.
	File remototypes.File `json:"file"`
}

// TeachFileResponse is the response object for TeachFile calls.
type TeachFileResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// TeachURLRequest is the request object for TeachURL calls.
type TeachURLRequest struct {

	// ID is an identifier describing the source, for example the filename.
	ID string `json:"id"`

	// Name is the name of the person in the image.
	Name string `json:"name"`

	// URL is the address of the image.
	URL string `json:"url"`
}

// TeachURLResponse is the response object for TeachURL calls.
type TeachURLResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// httpFaceboxServer is an internal type that provides an
// HTTP wrapper around Facebox.
type httpFaceboxServer struct {
	// service is the Facebox being exposed by this
	// server.
	service Facebox
	// server is the remotohttp.Server that this server is
	// registered with.
	server *remotohttp.Server
}

// handleCheckFaceprint is an http.Handler wrapper for Facebox.CheckFaceprint.
func (srv *httpFaceboxServer) handleCheckFaceprint(w http.ResponseWriter, r *http.Request) {
	var reqs []*CheckFaceprintRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]CheckFaceprintResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.CheckFaceprint(r.Context(), reqs[i])
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

// handleCheckFile is an http.Handler wrapper for Facebox.CheckFile.
func (srv *httpFaceboxServer) handleCheckFile(w http.ResponseWriter, r *http.Request) {
	var reqs []*CheckFileRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]CheckFileResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.CheckFile(r.Context(), reqs[i])
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

// handleCheckURL is an http.Handler wrapper for Facebox.CheckURL.
func (srv *httpFaceboxServer) handleCheckURL(w http.ResponseWriter, r *http.Request) {
	var reqs []*CheckURLRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]CheckURLResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.CheckURL(r.Context(), reqs[i])
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

// handleFaceprintCompare is an http.Handler wrapper for Facebox.FaceprintCompare.
func (srv *httpFaceboxServer) handleFaceprintCompare(w http.ResponseWriter, r *http.Request) {
	var reqs []*FaceprintCompareRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]FaceprintCompareResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.FaceprintCompare(r.Context(), reqs[i])
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

// handleGetState is an http.Handler wrapper for Facebox.GetState.
func (srv *httpFaceboxServer) handleGetState(w http.ResponseWriter, r *http.Request) {
	var reqs []*GetStateRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	// single file response

	if len(reqs) != 1 {
		if err := remotohttp.EncodeErr(w, r, errors.New("only single requests supported for file response endpoints")); err != nil {
			srv.server.OnErr(w, r, err)
			return
		}
		return
	}

	resp, err := srv.service.GetState(r.Context(), reqs[0])
	if err != nil {
		resp.Error = err.Error()
		if err := remotohttp.Encode(w, r, http.StatusOK, []interface{}{resp}); err != nil {
			srv.server.OnErr(w, r, err)
			return
		}
	}
	if resp.ContentType == "" {
		resp.ContentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", resp.ContentType)
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.QuoteToASCII(resp.Filename))
	if resp.ContentLength > 0 {
		w.Header().Set("Content-Length", strconv.Itoa(resp.ContentLength))
	}
	if _, err := io.Copy(w, resp.Data); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handlePutState is an http.Handler wrapper for Facebox.PutState.
func (srv *httpFaceboxServer) handlePutState(w http.ResponseWriter, r *http.Request) {
	var reqs []*PutStateRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]PutStateResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.PutState(r.Context(), reqs[i])
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

// handleRemoveID is an http.Handler wrapper for Facebox.RemoveID.
func (srv *httpFaceboxServer) handleRemoveID(w http.ResponseWriter, r *http.Request) {
	var reqs []*RemoveIDRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]RemoveIDResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.RemoveID(r.Context(), reqs[i])
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

// handleRename is an http.Handler wrapper for Facebox.Rename.
func (srv *httpFaceboxServer) handleRename(w http.ResponseWriter, r *http.Request) {
	var reqs []*RenameRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]RenameResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.Rename(r.Context(), reqs[i])
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

// handleRenameID is an http.Handler wrapper for Facebox.RenameID.
func (srv *httpFaceboxServer) handleRenameID(w http.ResponseWriter, r *http.Request) {
	var reqs []*RenameIDRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]RenameIDResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.RenameID(r.Context(), reqs[i])
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

// handleSimilarFile is an http.Handler wrapper for Facebox.SimilarFile.
func (srv *httpFaceboxServer) handleSimilarFile(w http.ResponseWriter, r *http.Request) {
	var reqs []*SimilarFileRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]SimilarFileResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.SimilarFile(r.Context(), reqs[i])
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

// handleSimilarID is an http.Handler wrapper for Facebox.SimilarID.
func (srv *httpFaceboxServer) handleSimilarID(w http.ResponseWriter, r *http.Request) {
	var reqs []*SimilarIDRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]SimilarIDResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.SimilarID(r.Context(), reqs[i])
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

// handleSimilarURL is an http.Handler wrapper for Facebox.SimilarURL.
func (srv *httpFaceboxServer) handleSimilarURL(w http.ResponseWriter, r *http.Request) {
	var reqs []*SimilarURLRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]SimilarURLResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.SimilarURL(r.Context(), reqs[i])
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

// handleTeachFaceprint is an http.Handler wrapper for Facebox.TeachFaceprint.
func (srv *httpFaceboxServer) handleTeachFaceprint(w http.ResponseWriter, r *http.Request) {
	var reqs []*TeachFaceprintRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]TeachFaceprintResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.TeachFaceprint(r.Context(), reqs[i])
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

// handleTeachFile is an http.Handler wrapper for Facebox.TeachFile.
func (srv *httpFaceboxServer) handleTeachFile(w http.ResponseWriter, r *http.Request) {
	var reqs []*TeachFileRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]TeachFileResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.TeachFile(r.Context(), reqs[i])
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

// handleTeachURL is an http.Handler wrapper for Facebox.TeachURL.
func (srv *httpFaceboxServer) handleTeachURL(w http.ResponseWriter, r *http.Request) {
	var reqs []*TeachURLRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]TeachURLResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.TeachURL(r.Context(), reqs[i])
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

// this is here so we don't get a compiler complaints.
func init() {
	var _ = remototypes.File{}
	var _ = strconv.Itoa(0)
	var _ = io.EOF
}
