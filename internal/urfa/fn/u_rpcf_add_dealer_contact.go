package fn

// rpcf_add_dealer_contact
func x13012(c conn, p Dict) Dict {
	c.Call(0x13012)
	putI(c, p, "id")
	putI(c, p, "user_id")
	putS(c, p, "person")
	putS(c, p, "descr")
	putS(c, p, "contact")
	putS(c, p, "email")
	putI(c, p, "email_notify")
	putS(c, p, "short_name")
	putS(c, p, "birthday")
	putI(c, p, "id_exec_man")
	c.Send()

	return nil
}
