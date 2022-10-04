package fn

// rpcf_user5_change_lifestream_credentials
func xU15054(c conn, p Dict) Dict {
	c.Call(-0x15054)
	putS(c, p, "lifestream_login")
	putS(c, p, "lifestream_password")
	putS(c, p, "lifestream_email")
	c.Send()

	return Dict{
		"lifestream_id": c.GetS(),
	}
}
