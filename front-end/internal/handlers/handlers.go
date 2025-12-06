package handlers

import (
	"fmt"
	"frontend/internal/config"
	"frontend/internal/models"
	"frontend/internal/render"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// ChatRoom is the handler for the chatroom page
func (m *Repository) ChatRoom(w http.ResponseWriter, r *http.Request) {
	// Get connected users
	resp, err := m.App.GRPCClient.GetAllConnectedusers(r.Context(), &emptypb.Empty{})
	if err != nil {
		fmt.Println("something break", err)
		fmt.Fprint(w, err)
		return
	}
	//Get all messages
	fmt.Println("getting message from grpc")
	response, err := m.App.GRPCClient.GetAllChatMessages(r.Context(), &emptypb.Empty{})
	if err != nil {
		fmt.Println("something break", err)
		fmt.Fprint(w, err)
		return
	}
	//Set Template data
	data := make(map[string]any)
	data["messsages"] = response.Result
	data["users"] = resp.Result
	// Render template
	render.RenderTemplate(w, r, "chatroom.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
