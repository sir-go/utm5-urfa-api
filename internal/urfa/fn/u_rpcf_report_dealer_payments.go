package fn

// rpcf_report_dealer_payments
func x1300f(c conn, p Dict) Dict {
	c.Call(0x1300f)
	putI(c, p, "dealer_id")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	return Dict{
		"payments_count": c.GetI(),
		"payments_sum":   c.GetD(),
	}
}
