package fn

// rpcf_block_card
func x4204(c conn, p Dict) Dict {
	c.Call(0x4204)
	putI(c, p, "card_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
