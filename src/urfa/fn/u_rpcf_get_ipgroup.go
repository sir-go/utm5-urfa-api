package fn

// rpcf_get_ipgroup
func x2902(c conn, p Dict) Dict {
	c.Call(0x2902)
	putI(c, p, "ipgroup_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2902" name="rpcf_get_ipgroup">
    <input>
      <integer name="ipgroup_id" />
    </input>
    <output>
      <string name="name" />
      <integer name="ipzone_count" />
      <for count="ipzone_count" from="0" name="i">
        <ip_address array_index="i" name="ip" />
        <ip_address array_index="i" name="mask" />
        <ip_address array_index="i" name="gateaway" />
      </for>
    </output>
  </function>


*/
