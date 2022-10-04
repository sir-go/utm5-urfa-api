package fn

// rpcf_add_user_log
func x1510c(c conn, p Dict) Dict {
	c.Call(0x1510c)
	putI(c, p, "user_id")
	putI(c, p, "what", 0)
	putS(c, p, "comment")
	c.Send()

	return nil
}
