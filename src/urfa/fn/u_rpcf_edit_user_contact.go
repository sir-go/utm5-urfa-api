package fn

// rpcf_edit_user_contact
func x2043(c conn, p Dict) Dict {
	c.Call(0x2043)
	putI(c, p, "cid")
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
