package fn

// rpcf_add_message
func x5001(c conn, p Dict) Dict {
	c.Call(0x5001)
	putI(c, p, "receiver_id")
	putS(c, p, "subject")
	putS(c, p, "message")
	putS(c, p, "mime")
	putI(c, p, "is_for_all")
	c.Send()

	return nil
}
