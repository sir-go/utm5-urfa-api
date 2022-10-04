package fn

// rpcf_remove_message_flag
func x5031(c conn, p Dict) Dict {
	c.Call(0x5031)
	putI(c, p, "message_id")
	putI(c, p, "flag")
	c.Send()

	return nil
}
