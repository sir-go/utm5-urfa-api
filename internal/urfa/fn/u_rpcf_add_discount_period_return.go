package fn

// rpcf_add_discount_period_return
func x2608(c conn, p Dict) Dict {
	c.Call(0x2608)
	putI(c, p, "static_id", 0)
	putI(c, p, "start_date")
	putI(c, p, "expire_date", 0)
	putI(c, p, "periodic_type", 3)
	putI(c, p, "custom_duration", 86400)
	putI(c, p, "discount_interval", 7)
	putI(c, p, "invoice_month", 0)
	c.Send()

	return Dict{
		"discount_period_id": c.GetI(),
	}
}
