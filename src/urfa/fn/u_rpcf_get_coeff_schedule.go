package fn

// rpcf_get_coeff_schedule
func x4705(c conn, p Dict) Dict {
	c.Call(0x4705)
	putI(c, p, "id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4705" name="rpcf_get_coeff_schedule">
		<input>
            <integer name="id" />
		</input>
		<output>
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
		</output>
	</function>


*/
