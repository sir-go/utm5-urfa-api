package fn

// rpcf_del_bank
func x6003(c conn, p Dict) Dict {
	c.Call(0x6003)
	putI(c, p, "bank_id")
	putS(c, p, "bic")
	c.Send()

	return nil
}
