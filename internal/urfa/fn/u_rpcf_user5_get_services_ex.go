package fn

// rpcf_user5_get_services_ex
func xU402f(c conn, _ Dict) Dict {
	c.Call(-0x402f)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x402f" name="rpcf_user5_get_services_ex">
    <input />
    <output>
      <integer name="links_size" />
      <for count="links_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="service_id" />
        <integer array_index="i" name="service_type" />
        <string array_index="i" name="service_name" />
        <string array_index="i" name="tariff_name" />
        <integer array_index="i" name="discount_period_start" />
        <integer array_index="i" name="discount_period_end" />
        <double array_index="i" name="cost" />
        <double array_index="i" name="discounted_in_curr_period" />
      </for>
    </output>
  </function>


*/
