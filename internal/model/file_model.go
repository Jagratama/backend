package model

type File struct {
	ID          int64  `json:"id"`
	FileName    string `json:"file_name"`
	FilePath    string `json:"file_path"`
	ContentType string `json:"content_type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
