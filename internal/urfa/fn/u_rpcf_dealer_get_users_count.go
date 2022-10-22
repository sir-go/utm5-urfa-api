package fn

// rpcf_dealer_get_users_count
func x7000000a(c conn, _ Dict) Dict {
	c.Call(0x7000000a)

	return Dict{
		"users_count": c.GetI(),
	}
}
