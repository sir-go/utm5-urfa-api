package fn

// rpcf_edit_coeff_schedule
func x4708(c conn, p Dict) Dict {
	c.Call(0x4708)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"id": c.GetI(),
	}
}

/* xml:
<function id="0x4708" name="rpcf_edit_coeff_schedule">
		<input>
            <integer name="id" />
            <integer name="create_date" />
            <integer name="change_date" />
            <integer name="items_size" />
          <for count="items_size" from="0" name="i">
            <integer array_index="i" name="schedule_id" />
            <integer array_index="i" name="start_date" />
            <integer array_index="i" name="end_date" />
            <double array_index="i" name="coeff" />
          </for>
		</input>
		<output>
            <integer name="id" />
		</output>
	</function>


*/
