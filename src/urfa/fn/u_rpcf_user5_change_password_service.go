package fn

// rpcf_user5_change_password_service
func xU4302(c conn, p Dict) Dict {
	c.Call(-0x4302)
	putI(c, p, "slink_id")
	putI(c, p, "item_id")
	putS(c, p, "login")
	putS(c, p, "old_password")
	putS(c, p, "new_password")
	putS(c, p, "new_password_ret")
	c.Send()

	return Dict{
		"status": c.GetI(),
	}
}
