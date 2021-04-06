package node_client

// Session is the type needed for creating/deleting sessions on the node
type Session struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetBalancesBodyResponse struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

type CreateJobBodyResponse struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}
