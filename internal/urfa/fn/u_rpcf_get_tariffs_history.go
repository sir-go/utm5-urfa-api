package fn

// rpcf_get_tariffs_history
func x301c(c conn, p Dict) Dict {
	c.Call(0x301c)
	putI(c, p, "aid")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x301c" name="rpcf_get_tariffs_history">
    <input>
      <integer name="aid" />
    </input>
    <output>
      <integer name="th_count" />
      <for count="th_count" from="0" name="i">
        <integer array_index="i" name="tariff_id" />
        <integer array_index="i" name="link_date" />
        <integer array_index="i" name="unlink_date" />
        <string array_index="i" name="tariff_name" />
      </for>
    </output>
  </function>


*/
