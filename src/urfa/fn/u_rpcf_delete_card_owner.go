package fn

// rpcf_delete_card_owner
func x4290(c conn, p Dict) Dict {
	c.Call(0x4290)
	putI(c, p, "pool_id")
	putI(c, p, "owned_id")
	c.Send()

	return nil
}
