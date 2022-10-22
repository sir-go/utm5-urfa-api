package fn

// rpcf_radius_disconnect_session
func x1071(c conn, p Dict) Dict {
	c.Call(0x1071)
	putS(c, p, "acct_session_id")
	putA(c, p, "nas_ip")
	c.Send()

	return Dict{
		"error_code": c.GetI(),
	}
}
