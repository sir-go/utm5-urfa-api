package fn

// rpcf_get_ipzone_ipv6
func x2804(c conn, p Dict) Dict {
	c.Call(0x2804)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"name": c.GetS(),
		"ip_zones": getMapOf(c, func() Dict {
			return Dict{
				"net":     c.GetA(),
				"mask":    c.GetI(),
				"gateway": c.GetA(),
			}
		}),
	}
}

/* xml:
<function id="0x2804" name="rpcf_get_ipzone_ipv6">
    <input>
      <integer name="id" />
    </input>
    <output>
      <string name="name" />
      <integer name="count" />
      <for count="count" from="0" name="i">
        <ip_address array_index="i" name="net" />
        <integer array_index="i" name="mask" />
        <ip_address array_index="i" name="gateaway" />
      </for>
    </output>
  </function>


*/
