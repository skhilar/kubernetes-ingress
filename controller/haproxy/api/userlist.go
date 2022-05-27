package api

func (c *haProxyClient) UserListExistsByGroup(group string) (exist bool, err error) {
	return false, nil
}

func (c *haProxyClient) UserListDeleteAll() (err error) {
	return nil
}

func (c *haProxyClient) UserListCreateByGroup(group string, userPasswordMap map[string][]byte) (err error) {
	return nil
}
