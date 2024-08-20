package closeio

func (c *Client) Update(newRecord CloseModel) ([]byte, error) {

	req, err := newRecord.Update(c.BaseURL)

	if err != nil {
		return nil, err
	}

	return c.makeHttpCall(req)
}
