package media

type MultimediaData struct {
	ContentType string `db:"content_type" json:"contentType"`
	Size        int    `db:"attachment_size" json:"size"`
	Data        []byte `db:"attachment_data" json:"data"`
}
