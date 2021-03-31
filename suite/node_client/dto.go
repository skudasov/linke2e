package node_client

// Session is the type needed for creating/deleting sessions on the node
type Session struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	Type          string            `json:"type"`
	Confirmations int               `json:"confirmations,omitempty"`
	Params        map[string]string `json:"params"`
}

type Initiator struct {
	Type   string            `json:"type"`
	Params map[string]string `json:"params"`
}

type CreateJobBody struct {
	Initiators []Initiator `json:"initiators"`
	Tasks      []Task      `json:"tasks"`
	Startat    string      `json:"startAt,omitempty"`
	Endat      interface{} `json:"endAt"`
	Minpayment string      `json:"minPayment"`
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
