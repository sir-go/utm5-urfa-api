package fn

// rpcf_edit_iptraffic_service_ex
func x1511(c conn, p Dict) Dict {
	c.Call(0x1511)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"service_id": c.GetI(),
	}
}

/* xml:
<function id="0x1511" name="rpcf_edit_iptraffic_service_ex">
    <input>
      <integer name="service_id" />
      <integer name="parent_id" />
      <integer name="tariff_id" />
      <string name="service_name" />
      <string name="comment" />
      <integer name="link_by_default" />
      <integer name="is_dynamic" />
      <integer name="contract_type" />
      <double name="cost" />
      <integer name="discount_method_t" />
      <integer name="sessions_limit" />
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
