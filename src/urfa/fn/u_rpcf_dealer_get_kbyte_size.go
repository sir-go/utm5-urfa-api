package fn

// rpcf_dealer_get_kbyte_size
func x7000004b(c conn, _ Dict) Dict {
	c.Call(0x7000004b)

	return Dict{
		"kbyte_size": c.GetI(),
	}
}
