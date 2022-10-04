package fn

// rpcf_get_prepaid_units
func x5500(c conn, p Dict) Dict {
	c.Call(0x5500)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5500" name="rpcf_get_prepaid_units">
    <input>
      <integer name="slink_id" />
    </input>
    <output>
      <integer name="bytes_in_mbyte" />
      <integer name="pinfo_size" />
      <for count="pinfo_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <long array_index="i" name="old" />
        <long array_index="i" name="cur" />
      </for>
    </output>
  </function>


*/
