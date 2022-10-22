package fn

// rpcf_remove_account
func x2034(c conn, p Dict) Dict {
	c.Call(0x2034)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"ret_code": c.GetI(),
	}
}
