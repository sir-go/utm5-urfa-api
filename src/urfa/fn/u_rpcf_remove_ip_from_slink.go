package fn

// rpcf_remove_ip_from_slink
func x1510e(c conn, p Dict) Dict {
	c.Call(0x1510e)
	putI(c, p, "slink_id")
	putA(c, p, "ip")
	putI(c, p, "mask")
	putS(c, p, "login")
	putS(c, p, "mac")
	putI(c, p, "switch_id", 0)
	putI(c, p, "port_id", 0)
	putI(c, p, "vlan_id", 0)
	c.Send()

	return Dict{
		"result":   c.GetI(),
		"err_desc": c.GetS(),
	}
}
