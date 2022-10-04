package fn

// rpcf_block_account
func x2037(c conn, p Dict) Dict {
	c.Call(0x2037)
	putI(c, p, "account_id")
	putI(c, p, "is_blocked")
	c.Send()

	return nil
}
