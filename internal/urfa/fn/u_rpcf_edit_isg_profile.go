package fn

// rpcf_edit_isg_profile
func x1401(c conn, p Dict) Dict {
	c.Call(0x1401)
	putI(c, p, "id")
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
