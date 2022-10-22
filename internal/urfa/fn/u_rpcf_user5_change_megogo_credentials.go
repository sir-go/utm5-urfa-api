package fn

// rpcf_user5_change_megogo_credentials
func xU15045(c conn, p Dict) Dict {
	c.Call(-0x15045)
	putS(c, p, "megogo_login")
	putS(c, p, "megogo_password")
	putS(c, p, "megogo_email")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
