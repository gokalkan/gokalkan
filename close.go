package gokalkan

func (cli *Client) Close() (err error) {
	err = cli.kc.Close()
	if err != nil {
		cli.log.Error("kc close error: ", err)
		return err
	}

	return nil
}
