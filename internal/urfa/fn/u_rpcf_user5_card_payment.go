package fn

// rpcf_user5_card_payment
func xU4205(c conn, p Dict) Dict {
	c.Call(-0x4205)
	putI(c, p, "account_id")
	putI(c, p, "card_id")
	putS(c, p, "secret")
	c.Send()

	return nil
}
