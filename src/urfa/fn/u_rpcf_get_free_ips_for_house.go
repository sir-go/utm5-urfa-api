package fn

// rpcf_get_free_ips_for_house
func x15101(c conn, p Dict) Dict {
	c.Call(0x15101)
	putI(c, p, "house_id")
	c.Send()

	res := Dict{"ips": getMapOf(c, func() Dict {
		return Dict{
			"ips_ip":    c.GetA(),
			"mask":      c.GetI(),
			"zone_name": c.GetS(),
		}
	})}
	res["error"] = c.GetS()

	return res
}

/* xml:
<function id="0x15101" name="rpcf_get_free_ips_for_house">
    <input>
      <integer name="house_id" />
    </input>
    <output>
      <integer name="ips_size" />
      <for count="ips_size" from="0" name="i">
        <ip_address array_index="i" name="ips_ip" />
        <integer array_index="i" name="mask" />
        <string array_index="i" name="zone_name" />
      </for>
      <string name="error" />
    </output>
  </function>


*/
