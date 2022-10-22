package fn

// rpcf_move_expired_cards
func x4206(c conn, _ Dict) Dict {
	c.Call(0x4206)

	return Dict{
		"result": c.GetI(),
	}
}
