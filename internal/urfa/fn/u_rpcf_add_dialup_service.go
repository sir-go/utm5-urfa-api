package fn

// rpcf_add_dialup_service
func x220d(c conn, p Dict) Dict {
	c.Call(0x220d)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x220d" name="rpcf_add_dialup_service">
    <input>
      <integer name="parent_id" />
      <integer name="tariff_id" />
      <integer name="service_id" />
      <string name="service_name" />
      <string name="comment" />
      <integer name="link_by_default" />
      <integer name="is_dynamic" />
      <integer name="contract_type" />
      <integer default="0" name="periodic_param" />
      <integer name="discount_method" />
      <double name="cost" />
      <integer name="start_date" />
      <integer name="expire_date" />
      <string name="pool_name" />
      <integer name="max_timeout" />
      <integer name="null_service_prepaid" />
      <integer name="radius_sessions_limit" />
      <string name="login_prefix" />
      <integer default="size(range_id)" name="cost_size" />
      <for count="size(range_id)" from="0" name="i">
        <double array_index="i" name="range_cost" />
        <integer array_index="i" name="range_id" />
      </for>
    </input>
    <output>
      <integer name="result" />
    </output>
  </function>


*/
