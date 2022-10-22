package fn

// rpcf_add_message_new
func x500d(c conn, p Dict) Dict {
	c.Call(0x500d)
	putI(c, p, "receiver_id")
	putI(c, p, "receiver_type")
	putS(c, p, "subject")
	putS(c, p, "message")
	putS(c, p, "mime")
	c.Send()

	return nil
}
