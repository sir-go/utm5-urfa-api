package fn

// rpcf_add_user_contact
func x2042(c conn, p Dict) Dict {
	c.Call(0x2042)
	putI(c, p, "user_id")
	putS(c, p, "descr")
	putS(c, p, "reason")
	putS(c, p, "person")
	putS(c, p, "short_name")
	putS(c, p, "contact")
	putS(c, p, "email")
	putI(c, p, "id_exec_man")
	c.Send()

	return nil
}
