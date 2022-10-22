package fn

// rpcf_add_payment_for_account_notify
func x3113(c conn, p Dict) Dict {
	c.Call(0x3113)
	putI(c, p, "account_id")
	putI(c, p, "param")
	putD(c, p, "payment_incurrency")
	putI(c, p, "currency_id")
	putI(c, p, "actual_date")
	putI(c, p, "burn_time")
	putI(c, p, "method")
	putS(c, p, "admin_comment")
	putS(c, p, "comment")
	putS(c, p, "payment_ext_number")
	putI(c, p, "payment_to_invoice")
	putI(c, p, "turn_on_inet")
	putI(c, p, "notify")
	putS(c, p, "hash")
	c.Send()

	return Dict{
		"payment_transaction_id": c.GetI(),
	}
}
