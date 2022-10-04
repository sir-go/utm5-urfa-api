package fn

// rpcf_user5_get_doc_data
func xU15032(c conn, p Dict) Dict {
	c.Call(-0x15032)
	putI(c, p, "invoice_id")
	c.Send()

	return Dict{
		"supp_name":            c.GetS(),
		"supp_short_name":      c.GetS(),
		"supp_jur_address":     c.GetS(),
		"supp_acct_address":    c.GetS(),
		"supp_inn":             c.GetS(),
		"supp_kpp":             c.GetS(),
		"bank_name":            c.GetS(),
		"bank_city":            c.GetS(),
		"bank_bic":             c.GetS(),
		"bank_kschet":          c.GetS(),
		"supp_account":         c.GetS(),
		"supp_headman":         c.GetS(),
		"supp_bookeeper":       c.GetS(),
		"supp_short_headman":   c.GetS(),
		"supp_short_bookeeper": c.GetS(),
	}
}
