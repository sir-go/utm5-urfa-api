package fn

// rpcf_verify_dealer_cache
func x1301d(c conn, _ Dict) Dict {
	c.Call(0x1301d)

	return Dict{
		"result": c.GetI(),
	}
}
