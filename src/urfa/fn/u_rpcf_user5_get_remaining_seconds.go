package fn

// rpcf_user5_get_remaining_seconds
func xU2027(c conn, p Dict) Dict {
	c.Call(-0x2027)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"remaining_seconds":  c.GetI(),
		"downloaded_seconds": c.GetI(),
	}
}
