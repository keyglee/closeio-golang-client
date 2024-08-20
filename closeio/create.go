package closeio

func (c *Client) Create(newRecord CloseModel) ([]byte, error) {

	req, err := newRecord.Create(c.BaseURL)

	if err != nil {
		return nil, err
	}

	return c.makeHttpCall(req)
}
