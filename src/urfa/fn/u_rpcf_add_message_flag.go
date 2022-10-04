package fn

// rpcf_add_message_flag
func x5030(c conn, p Dict) Dict {
	c.Call(0x5030)
	putI(c, p, "message_id")
	putI(c, p, "flag")
	c.Send()

	return nil
}
