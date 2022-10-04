package fn

// rpcf_user5_add_mime_message
func xU4034(c conn, p Dict) Dict {
	c.Call(-0x4034)
	putS(c, p, "subject")
	putS(c, p, "message")
	putS(c, p, "mime")
	putI(c, p, "state")
	c.Send()

	return nil
}
