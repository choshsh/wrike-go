package wrike

type Attachments struct {
	Kind string       `json:"kind"`
	Data []Attachment `json:"data"`
}

type Attachment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (w *Wrike) Attachments(taskId string) Attachments {
	urlQuery := map[string]string{
		"withUrls": "true",
	}
	attachments := Attachments{}
	w.newAPI("/tasks/"+taskId+"/attachments", urlQuery, &attachments)

	return attachments
}
