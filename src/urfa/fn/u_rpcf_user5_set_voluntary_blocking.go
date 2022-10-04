package fn

// rpcf_user5_set_voluntary_blocking
func xU15015(c conn, p Dict) Dict {
	c.Call(-0x15015)
	putI(c, p, "account_id", 0)
	putI(c, p, "block_start")
	putI(c, p, "block_end")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
