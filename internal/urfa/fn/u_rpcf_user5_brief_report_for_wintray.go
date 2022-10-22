package fn

// rpcf_user5_brief_report_for_wintray
func xU4026(c conn, _ Dict) Dict {
	c.Call(-0x4026)

	return Dict{
		"user_int_status": c.GetI(),
		"balance":         c.GetD(),
	}
}
