package closeio

func (c *Client) Delete(newRecord CloseModel) ([]byte, error) {

	req, err := newRecord.Delete(c.BaseURL)

	if err != nil {
		return nil, err
	}

	return c.makeHttpCall(req)
}
