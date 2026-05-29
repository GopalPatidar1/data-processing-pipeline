package routes

import (
	"encoding/json"
	"net/http"
)

type File struct {
	ID       int    `json:"id"`
	FileName string `json:"file_name"`
	Size     string `json:"size"`
}

func getFiles(w http.ResponseWriter, r *http.Request) {

	files := []File{
		{
			ID:       1,
			FileName: "resume.pdf",
			Size:     "2MB",
		},
		{
			ID:       2,
			FileName: "image.png",
			Size:     "5MB",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(files)
}

func FileRoutes() {
	http.HandleFunc("/files", getFiles)
}