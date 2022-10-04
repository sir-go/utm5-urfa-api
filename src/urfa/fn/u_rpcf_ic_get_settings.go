package fn

// rpcf_ic_get_settings
func x14001(c conn, _ Dict) Dict {
	c.Call(0x14001)

	return Dict{
		"send_users":          c.GetI(),
		"send_inv":            c.GetI(),
		"send_payments":       c.GetI(),
		"recv_payments":       c.GetI(),
		"sync_card":           c.GetI(),
		"sync_empty_name":     c.GetI(),
		"sync_empty_inn":      c.GetI(),
		"sync_empty_kpp":      c.GetI(),
		"sync_users_since":    c.GetI(),
		"sync_users_till":     c.GetI(),
		"sync_inv_since":      c.GetI(),
		"sync_inv_till":       c.GetI(),
		"sync_payments_since": c.GetI(),
		"sync_payments_till":  c.GetI(),
	}
}
