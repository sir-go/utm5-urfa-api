package fn

// rpcf_user5_get_user_megogo_link
func xU15043(c conn, _ Dict) Dict {
	c.Call(-0x15043)

	return Dict{
		"link": c.GetS(),
	}
}
