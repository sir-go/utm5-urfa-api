package fn

// rpcf_get_bytes_in_kb
func x10002(c conn, _ Dict) Dict {
	c.Call(0x10002)

	return Dict{
		"bytes_in_kb": c.GetI(),
	}
}
