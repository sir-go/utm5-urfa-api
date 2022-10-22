package fn

// rpcf_refresh_archives
func x00fb(c conn, _ Dict) Dict {
	c.Call(0x00fb)

	return Dict{
		"result": c.GetI(),
	}
}
