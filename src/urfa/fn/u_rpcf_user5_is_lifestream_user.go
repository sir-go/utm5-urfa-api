package fn

// rpcf_user5_is_lifestream_user
func xU15055(c conn, _ Dict) Dict {
	c.Call(-0x15055)

	return Dict{
		"result": c.GetI(),
	}
}
