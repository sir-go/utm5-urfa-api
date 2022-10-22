package fn

// rpcf_get_switches_count
func x1155(c conn, _ Dict) Dict {
	c.Call(0x1155)

	return Dict{
		"count": c.GetI(),
	}
}
