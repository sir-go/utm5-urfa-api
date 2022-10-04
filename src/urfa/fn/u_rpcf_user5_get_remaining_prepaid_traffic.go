package fn

// rpcf_user5_get_remaining_prepaid_traffic
func xU4038(c conn, _ Dict) Dict {
	c.Call(-0x4038)

	return Dict{
		"result": c.GetI(),
	}
}
