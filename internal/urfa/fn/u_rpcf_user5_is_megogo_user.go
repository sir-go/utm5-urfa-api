package fn

// rpcf_user5_is_megogo_user
func xU15046(c conn, _ Dict) Dict {
	c.Call(-0x15046)

	return Dict{
		"result": c.GetI(),
	}
}
