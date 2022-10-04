package fn

// rpcf_user5_get_services
func xU4023(c conn, _ Dict) Dict {
	c.Call(-0x4023)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4023" name="rpcf_user5_get_services">
    <input />
    <output>
      <integer name="links_size" />
      <for count="links_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="service_id" />
        <integer array_index="i" name="service_type" />
        <string array_index="i" name="service_name" />
        <string array_index="i" name="tariff_name" />
        <string array_index="i" name="discount_period" />
        <double array_index="i" name="cost" />
        <double array_index="i" name="discounted_in_curr_period" />
      </for>
    </output>
  </function>


*/
