package fn

// rpcf_archivate
func x00fe(c conn, _ Dict) Dict {
	c.Call(0x00fe)

	return Dict{
		"result": c.GetI(),
	}
}
