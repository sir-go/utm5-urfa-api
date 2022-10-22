package fn

// rpcf_get_supplier
func x8013(c conn, p Dict) Dict {
	c.Call(0x8013)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":              c.GetI(),
		"name":            c.GetS(),
		"short_name":      c.GetS(),
		"jur_adress":      c.GetS(),
		"act_adress":      c.GetS(),
		"inn":             c.GetS(),
		"kpp":             c.GetS(),
		"bank_id":         c.GetI(),
		"account":         c.GetS(),
		"headman":         c.GetS(),
		"bookeeper":       c.GetS(),
		"short_headman":   c.GetS(),
		"short_bookeeper": c.GetS(),
		"supp_balance":    c.GetD(),
		"tax_rate":        c.GetD(),
		"type":            c.GetI(),
	}
}
