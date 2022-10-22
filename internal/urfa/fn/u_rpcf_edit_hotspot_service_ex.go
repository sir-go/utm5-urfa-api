package fn

// rpcf_edit_hotspot_service_ex
func x1507(c conn, p Dict) Dict {
	c.Call(0x1507)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"service_id": c.GetI(),
	}
}

/* xml:
<function id="0x1507" name="rpcf_edit_hotspot_service_ex">
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
      <integer name="radius_sessions_limit" />
      <double name="recv_cost" />
      <string name="rate_limit" />
      <integer name="allowed_net_size" />
      <for count="allowed_net_size" from="0" name="i">
        <ip_address array_index="i" name="allowed_net_ip" />
        <integer array_index="i" name="allowed_net_value" />
      </for>
      <integer name="periodic_service_size" />
      <for count="periodic_service_size" from="0" name="i">
        <double array_index="i" name="periodic_service_cost" />
        <integer array_index="i" name="periodic_service_id" />
      </for>
    </input>
    <output>
        <integer name="service_id" />
    </output>
  </function>


*/
