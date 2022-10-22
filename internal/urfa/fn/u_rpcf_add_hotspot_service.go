package fn

// rpcf_add_hotspot_service
func x2208(c conn, p Dict) Dict {
	c.Call(0x2208)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"sid": c.GetI(),
	}
}

/* xml:
<function id="0x2208" name="rpcf_add_hotspot_service">
    <input>
      <integer name="fictive" />
      <integer name="tariff_id" />
      <integer name="service_id" />
      <string name="service_name" />
      <string name="comment" />
      <integer name="link_by_default" />
      <integer name="is_dynamic" />
      <integer name="contract_type" />
      <integer name="param1" />
      <integer name="discount_method" />
      <double name="cost" />
      <integer name="start_date" />
      <integer name="expire_date" />
      <integer name="null_service_prepaid" />
      <integer name="radius_sessions_limit" />
      <double name="recv_cost" />
      <string name="rate_limit" />
      <integer name="allowed_net_size" />
      <for count="allowed_net_size" from="0" name="i">
        <ip_address array_index="i" name="ip" />
        <integer array_index="i" name="value1" />
      </for>
      <integer name="periodic_service_size" />
      <for count="periodic_service_size" from="0" name="i">
        <double array_index="i" name="cost_p" />
        <integer array_index="i" name="id_p" />
      </for>
    </input>
    <output>
        <integer name="sid" />
    </output>
  </function>


*/
