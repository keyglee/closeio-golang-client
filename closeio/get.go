package closeio

func (c *Client) Get(newRecord CloseModel) ([]byte, error) {

	req, err := newRecord.Read(c.BaseURL)
	if err != nil {
		return nil, err
	}

	return c.makeHttpCall(req)
}
