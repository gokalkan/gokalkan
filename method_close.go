package kalkan

// Close закрывает связь с динамической библиотекой
func (cli *Client) Close() error {
	cli.XMLFinalize()
	cli.Finalize()
	return cli.handler.Close()
}
