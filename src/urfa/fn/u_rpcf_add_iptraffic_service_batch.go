package fn

// rpcf_add_iptraffic_service_batch
func x2121(c conn, p Dict) Dict {
	c.Call(0x2121)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"service_id": c.GetI(),
	}
}

/* xml:
<function id="0x2121" name="rpcf_add_iptraffic_service_batch">
    <input>
      <integer name="parent_id" />
      <integer name="tariff_id" />
      <integer name="service_id" />
      <string name="service_name" />
      <string name="comment" />
      <integer name="link_by_default" />
      <integer name="is_dynamic" />
      <double name="cost" />
      <integer name="parameter" />
      <integer name="discount_method_t" />
      <integer name="start_date" />
      <integer name="expire_date" />
      <integer name="null_service_prepaid" />
      <integer default="size(tclass_b)" name="num_of_borders" />
      <for count="size(tclass_b)" from="0" name="i">
        <integer array_index="i" name="tclass_b" />
        <long array_index="i" name="size_b" />
        <double array_index="i" name="cost_b" />
      </for>
      <integer default="size(tclass_p)" name="num_of_prepaid" />
      <for count="size(tclass_p)" from="0" name="i">
        <integer array_index="i" name="tclass_p" />
        <long array_index="i" name="size_p" />
        <long array_index="i" name="size_max_p" />
      </for>
      <integer default="size(tcid)" name="num_of_groups" />
      <for count="size(tcid)" from="0" name="i">
        <integer array_index="i" name="tcid" />
        <integer array_index="i" name="gid" />
      </for>
    </input>
    <output>
      <integer name="service_id" />
    </output>
  </function>


*/
