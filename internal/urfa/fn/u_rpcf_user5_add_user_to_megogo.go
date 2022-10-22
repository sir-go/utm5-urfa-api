package fn

// rpcf_user5_add_user_to_megogo
func xU15044(c conn, p Dict) Dict {
	c.Call(-0x15044)
	putS(c, p, "megogo_login")
	putS(c, p, "megogo_password")
	putS(c, p, "megogo_email")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
