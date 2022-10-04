package fn

// rpcf_card_list
func x4200(c conn, p Dict) Dict {
	c.Call(0x4200)
	putI(c, p, "pool_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4200" name="rpcf_card_list">
    <input>
      <integer name="pool_id" />
    </input>
    <output>
      <integer name="cpi_owners_size" />
      <for count="cpi_owners_size" from="0" name="i">
        <integer array_index="i" name="owners" />
      </for>
      <integer name="info_size" />
      <integer name="time0" />
      <for count="info_size" from="0" name="i">
        <integer array_index="i" name="card_id" />
        <integer array_index="i" name="pool_id" />
        <string array_index="i" name="secret" />
        <double array_index="i" name="balance" />
        <integer array_index="i" name="currency" />
        <integer array_index="i" name="expire" />
        <integer array_index="i" name="days" />
        <integer array_index="i" name="is_used" />
        <integer array_index="i" name="tp_id" />
        <integer array_index="i" name="is_blocked" />
      </for>
    </output>
  </function>


*/
