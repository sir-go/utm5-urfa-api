package fn

// rpcf_user5_get_user_megogo_status
func xU15042(c conn, _ Dict) Dict {
	c.Call(-0x15042)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15042" name="rpcf_user5_get_user_megogo_status">
    <input>
    </input>
    <output>
      <integer name="user_tariffs_size" />
      <for count="user_tariffs_size" from="0" name="i">
          <integer array_index="i" name="tariff_plan_id" />
          <integer array_index="i" name="tariff_link_id" />
          <integer array_index="i" name="account_id" />
          <integer array_index="i" name="next_tariff_id" />
      </for>

    </output>
  </function>


*/
