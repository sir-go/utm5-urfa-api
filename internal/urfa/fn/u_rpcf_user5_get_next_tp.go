package fn

// rpcf_user5_get_next_tp
func xU15005(c conn, p Dict) Dict {
	c.Call(-0x15005)
	putI(c, p, "account_id", 0)
	putI(c, p, "tp_link", 0)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15005" name="rpcf_user5_get_next_tp">
    <input>
      <integer default="0" name="account_id" />
      <integer default="0" name="tp_link" />
    </input>
    <output>
      <integer name="user_tariffs_size" />
      <for count="user_tariffs_size" from="0" name="i">
        <integer array_index="i" name="tariff_id_array" />
        <string array_index="i" name="tariff_name_array" />
        <string array_index="i" name="tariff_comments_array" />
	<double array_index="i" name="min_balance" />
	<integer array_index="i" name="use_min_balance" />
	<double array_index="i" name="free_balance" />
	<integer array_index="i" name="use_free_balance" />
	<double array_index="i" name="cost" />
	<integer array_index="i" name="can_change" />
      </for>

	<double name="balance" />
    </output>
  </function>


*/
