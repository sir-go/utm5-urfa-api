package fn

// rpcf_get_isg_profile
func x1400(c conn, p Dict) Dict {
	c.Call(0x1400)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":                     c.GetI(),
		"name":                   c.GetS(),
		"login_type":             c.GetI(),
		"password_type":          c.GetI(),
		"blocked_account_code":   c.GetI(),
		"unblocked_account_code": c.GetI(),
		"lease_address":          c.GetI(),
		"static_password":        c.GetS(),
		"send_coa":               c.GetI(),
		"session_time":           c.GetI(),
	}
}
