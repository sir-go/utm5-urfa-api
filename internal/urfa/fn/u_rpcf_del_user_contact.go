package fn

// rpcf_del_user_contact
func x2023(c conn, p Dict) Dict {
	c.Call(0x2023)
	putI(c, p, "id")
	c.Send()

	return nil
}
