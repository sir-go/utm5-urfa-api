package fn

// rpcf_user5_get_invoice_data
func xU4022(c conn, _ Dict) Dict {
	c.Call(-0x4022)

	return Dict{
		"full_name":         c.GetS(),
		"actual_address":    c.GetS(),
		"juridical_address": c.GetS(),
		"basic_account":     c.GetI(),
		"payment_recv":      c.GetS(),
		"inn":               c.GetS(),
		"bank_account":      c.GetS(),
		"bank_name":         c.GetS(),
		"bank_city":         c.GetS(),
		"bank_bic":          c.GetS(),
		"bank_ks":           c.GetS(),
	}
}
