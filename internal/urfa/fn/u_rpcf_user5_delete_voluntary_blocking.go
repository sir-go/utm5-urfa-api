package fn

// rpcf_user5_delete_voluntary_blocking
func xU15016(c conn, p Dict) Dict {
	c.Call(-0x15016)
	putI(c, p, "account_id", 0)
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
