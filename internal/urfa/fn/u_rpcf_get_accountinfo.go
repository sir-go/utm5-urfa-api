package fn

// rpcf_get_accountinfo
func x15109(c conn, p Dict) Dict {
	c.Call(0x15109)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"is_blocked":       c.GetI(),
		"vat_rate":         c.GetD(),
		"sale_tax_rate":    c.GetD(),
		"credit":           c.GetD(),
		"balance":          c.GetD(),
		"int_status":       c.GetI(),
		"unlimited":        c.GetI(),
		"auto_enable_inet": c.GetI(),
		"external_id":      c.GetS(),
	}
}
