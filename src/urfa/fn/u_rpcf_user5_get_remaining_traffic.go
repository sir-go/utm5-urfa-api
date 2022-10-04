package fn

// rpcf_user5_get_remaining_traffic
func xU2026(c conn, p Dict) Dict {
	c.Call(-0x2026)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"traffic_remaining_mb":  c.GetD(),
		"traffic_downloaded_mb": c.GetD(),
	}
}
