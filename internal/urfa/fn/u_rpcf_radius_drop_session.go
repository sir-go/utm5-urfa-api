package fn

// rpcf_radius_drop_session
func x1072(c conn, p Dict) Dict {
	c.Call(0x1072)
	putS(c, p, "acct_session_id")
	putA(c, p, "nas_ip")
	c.Send()

	return Dict{
		"error_code": c.GetI(),
	}
}
