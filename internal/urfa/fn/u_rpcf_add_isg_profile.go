package fn

// rpcf_add_isg_profile
func x1402(c conn, p Dict) Dict {
	c.Call(0x1402)
	putS(c, p, "name")
	putI(c, p, "login_type")
	putI(c, p, "password_type")
	putI(c, p, "blocked_account_code")
	putI(c, p, "unblocked_account_code")
	putI(c, p, "lease_address")
	putS(c, p, "static_password")
	putI(c, p, "send_coa")
	putI(c, p, "session_time")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
