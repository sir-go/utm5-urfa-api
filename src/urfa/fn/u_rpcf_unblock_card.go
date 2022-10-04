package fn

// rpcf_unblock_card
func x4205(c conn, p Dict) Dict {
	c.Call(0x4205)
	putI(c, p, "card_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
