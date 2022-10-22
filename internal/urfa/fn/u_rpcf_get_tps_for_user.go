package fn

// rpcf_get_tps_for_user
func x301a(c conn, p Dict) Dict {
	c.Call(0x301a)
	putI(c, p, "uid")
	putI(c, p, "aid")
	putI(c, p, "tpid")
	putI(c, p, "tplink")
	putI(c, p, "unused")
	c.Send()

	return Dict{
		"services": getMapOf(c, func() Dict {
			return Dict{
				"sid":          c.GetI(),
				"service_name": c.GetS(),
				"service_type": c.GetI(),
				"comment":      c.GetS(),
				"slink":        c.GetI(),
				"value":        c.GetI(),
			}
		}),
	}
}

/* xml:
<function id="0x301a" name="rpcf_get_tps_for_user">
    <input>
      <integer name="uid" />
      <integer name="aid" />
      <integer name="tpid" />
      <integer name="tplink" />
      <integer name="unused" />
    </input>
    <output>
      <integer name="service_size" />
      <for count="service_size" from="0" name="i">
        <integer array_index="i" name="sid" />
        <string array_index="i" name="service_name" />
        <integer array_index="i" name="service_type" />
        <string array_index="i" name="comment" />
        <integer array_index="i" name="slink" />
        <integer array_index="i" name="value" />
      </for>
    </output>
  </function>


*/
