package fn

// rpcf_get_timezone
func x11113(c conn, _ Dict) Dict {
	c.Call(0x11113)

	return Dict{
		"tzname": c.GetS(),
	}
}
