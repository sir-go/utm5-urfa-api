package fn

// rpcf_edit_dialup_service_ex
func x1509(c conn, p Dict) Dict {
	c.Call(0x1509)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"service_id": c.GetI(),
	}
}

/* xml:
<function id="0x1509" name="rpcf_edit_dialup_service_ex">
    <input>
      <integer name="service_id" />
      <integer name="parent_id" />
      <integer name="tariff_id" />
      <string name="service_name" />
      <string name="comment" />
      <integer name="link_by_default" />
      <integer name="is_dynamic" />
      <integer name="contract_type" />
      <integer name="discount_method" />
      <double name="cost" />
      <string name="pool_name" />
      <integer name="max_timeout" />
      <integer name="radius_sessions_limit" />
      <string name="login_prefix" />
      <integer default="size(range_id)" name="cost_size" />
      <for count="size(range_id)" from="0" name="i">
        <double array_index="i" name="range_cost" />
        <integer array_index="i" name="range_id" />
      </for>
    </input>
    <output>
      <integer name="service_id" />
    </output>
  </function>


*/
