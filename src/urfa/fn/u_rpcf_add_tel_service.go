package fn

// rpcf_add_tel_service
func x1043d(c conn, p Dict) Dict {
	c.Call(0x1043d)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"service_id": c.GetI(),
	}
}

/* xml:
<function id="0x1043d" name="rpcf_add_tel_service">
    <input>
		<integer name="parent_id" />
	    <integer name="tariff_id" />
		<string name="service_name" />
		<string name="comment" />
		<integer name="link_by_default" />
		<integer name="is_dynamic" />
       <integer name="contract_type" />
		<double name="cost" />
		<integer name="discount_method" />
                <integer name="radius_sessions_limit" />
		<integer name="tariff_class_count" />
		<for count="tariff_class_count" from="0" name="i">
			<integer array_index="i" name="tariff_class" />
			<integer array_index="i" name="border_count" />
			<for count="border_count" from="0" name="j">
			    <integer array_index="i,j" name="border" />
				<double array_index="i,j" name="cost_mult" />
			</for>
			<integer array_index="i" name="timerange_count" />
			<for count="timerange_count" from="0" name="j">
				<integer array_index="i,j" name="timerange_id" />
				<double array_index="i,j" name="cost" />
			</for>
			<integer array_index="i" name="prepaid_array" />
			<double array_index="i" name="fixed_cost_array" />
		</for>
		<integer name="free_time" />
		<integer name="first_period_length" />
		<integer name="first_period_step" />
		<integer name="second_period_step" />
		<integer name="time_unit_size" />
		<integer name="discount_free_time" />
                <double name="minimum_charge" />
	</input>
	<output>
		<integer name="service_id" />
	</output>
  </function>


*/
