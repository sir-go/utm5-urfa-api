package fn

// rpcf_get_tariff
func x3011(c conn, p Dict) Dict {
	c.Call(0x3011)
	putI(c, p, "tariff_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x3011" name="rpcf_get_tariff">
    <input>
      <integer name="tariff_id" />
    </input>
    <output>
      <string name="tariff_name" />
      <integer name="tariff_create_date" />
      <integer name="who_create" />
      <string name="who_create_login" />
      <integer name="tariff_change_date" />
      <integer name="who_change" />
      <string name="who_change_login" />
      <integer name="tariff_expire_date" />
      <integer name="tariff_is_blocked" />
      <integer name="tariff_balance_rollover" />
      <integer name="services_count" />
      <for count="services_count" from="0" name="i">
        <integer array_index="i" name="service_id_array" />
        <integer array_index="i" name="service_type_array" />
        <string array_index="i" name="service_name_array" />
        <string array_index="i" name="comment_array" />
        <integer array_index="i" name="link_by_default_array" />
        <integer array_index="i" name="is_dynamic_array" />
      </for>
    </output>
  </function>


*/
