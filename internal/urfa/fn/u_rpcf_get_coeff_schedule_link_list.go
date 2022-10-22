package fn

// rpcf_get_coeff_schedule_link_list
func x4711(c conn, _ Dict) Dict {
	c.Call(0x4711)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4711" name="rpcf_get_coeff_schedule_link_list">
		<input>
		</input>
		<output>
            <integer name="list_size" />
          <for count="list_size" from="0" name="i">
            <integer array_index="i" name="id" />
            <integer array_index="i" name="scheme_id" />
            <integer array_index="i" name="slink_id" />
            <integer array_index="i" name="schedule_id" />
            <integer array_index="i" name="change_policy" />
            <integer array_index="i" name="create_date" />
            <integer array_index="i" name="change_date" />
          </for>
		</output>
	</function>


*/
