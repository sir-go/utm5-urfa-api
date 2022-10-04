package fn

// rpcf_verify_archives
func x00fc(c conn, _ Dict) Dict {
	c.Call(0x00fc)

	return Dict{
		"result": c.GetI(),
	}
}
