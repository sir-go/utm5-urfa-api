package fn

// rpcf_get_once_service_link_new
func x2556(c conn, p Dict) Dict {
	c.Call(0x2556)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"discount_date": c.GetI(),
		"quantity":      c.GetD(),
		"invoice_id":    c.GetI(),
		"cost_coef":     c.GetD(),
	}
}
