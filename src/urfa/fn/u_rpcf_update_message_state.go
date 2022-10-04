package fn

// rpcf_update_message_state
func x5029(c conn, p Dict) Dict {
	c.Call(0x5029)
	putI(c, p, "mid")
	putI(c, p, "state")
	c.Send()

	return nil
}
