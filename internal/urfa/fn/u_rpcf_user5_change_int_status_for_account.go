package fn

// rpcf_user5_change_int_status_for_account
func xU4049(c conn, p Dict) Dict {
	c.Call(-0x4049)
	putI(c, p, "account_id")
	putI(c, p, "int_status")
	c.Send()

	return nil
}
