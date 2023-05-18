package model

type Response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Vertex struct {
	Key          string `json:"key"`
	Value        string `json:"value"`
	UpdatedValue string `json:"updatedValue"`
}
