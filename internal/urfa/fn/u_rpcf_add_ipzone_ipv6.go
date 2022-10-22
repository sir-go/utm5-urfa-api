package fn

// rpcf_add_ipzone_ipv6
func x2803(c conn, p Dict) Dict {
	c.Call(0x2803)
	putI(c, p, "id")
	putS(c, p, "name")

	return Dict{
		"id": c.GetI(),
	}
}

/* xml:
<function id="0x2803" name="rpcf_add_ipzone_ipv6">
    <input>
      <integer name="id" />
      <string name="name" />
      <integer name="count" />
      <for count="count" from="0" name="i">
        <ip_address array_index="i" name="net" />
        <integer array_index="i" name="mask" />
        <ip_address array_index="i" name="gateaway" />
      </for>
    </input>
    <output>
      <integer name="id" />
    </output>
  </function>


*/
