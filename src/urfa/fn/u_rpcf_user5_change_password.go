package fn

// rpcf_user5_change_password
func xU4021(c conn, p Dict) Dict {
	c.Call(-0x4021)
	putS(c, p, "old_password")
	putS(c, p, "new_password")
	putS(c, p, "new_password_ret")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
