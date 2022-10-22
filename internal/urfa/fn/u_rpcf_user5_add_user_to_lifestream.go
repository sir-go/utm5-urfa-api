package fn

// rpcf_user5_add_user_to_lifestream
func xU15053(c conn, p Dict) Dict {
	c.Call(-0x15053)
	putS(c, p, "lifestream_login")
	putS(c, p, "lifestream_password")
	putS(c, p, "lifestream_email")
	c.Send()

	return Dict{
		"lifestream_id": c.GetS(),
	}
}
