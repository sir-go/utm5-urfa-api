package fn

// rpcf_user5_get_tlink
func xU15004(c conn, p Dict) Dict {
	c.Call(-0x15004)
	putI(c, p, "account_id", 0)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15004" name="rpcf_user5_get_tlink">
    <input>
      <integer default="0" name="account_id" />
    </input>
    <output>
      <integer name="user_tariffs_size" />
      <for count="user_tariffs_size" from="0" name="i">
        <integer array_index="i" name="tariff_link_id_array" />
        <integer array_index="i" name="tariff_current_array" />
        <string array_index="i" name="tariff_current_name_array" />
        <string array_index="i" name="tariff_current_comment_array" />
        <integer array_index="i" name="tariff_next_array" />
        <string array_index="i" name="tariff_next_name_array" />
        <string array_index="i" name="tariff_next_comment_array" />
        <integer array_index="i" name="discount_period_id_array" />
        <integer array_index="i" name="discount_period_start_array" />
        <integer array_index="i" name="discount_period_end_array" />
      </for>
    </output>
  </function>


*/
