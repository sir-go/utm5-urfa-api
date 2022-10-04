package fn

// rpcf_get_dialup_service_ex
func x2233(c conn, p Dict) Dict {
	c.Call(0x2233)
	putI(c, p, "sid")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2233" name="rpcf_get_dialup_service_ex">
    <input>
      <integer name="sid" />
    </input>
    <output>
      <string name="service_name" />
      <string name="comment" />
      <integer name="link_by_default" />
      <integer name="is_dynamic" />
      <integer name="contract_type" />
      <double name="cost" />
      <integer name="deprecated" />
      <integer name="discount_method" />
      <string name="pool_name" />
      <integer name="max_timeout" />
      <integer name="null_service_prepaid" />
      <integer name="radius_sessions_limit" />
      <string name="login_prefix" />
      <integer name="cost_size" />
      <for count="cost_size" from="0" name="i">
        <string array_index="i" name="tr_time" />
        <double array_index="i" name="param" />
        <integer array_index="i" name="id" />
      </for>
      <integer name="is_parent_id" />
      <integer name="tariff_id" />
      <integer name="parent_id" />
    </output>
  </function>


*/
