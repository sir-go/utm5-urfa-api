package fn

// rpcf_get_coeff_schedule_list
func x4706(c conn, _ Dict) Dict {
	c.Call(0x4706)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4706" name="rpcf_get_coeff_schedule_list">
		<input>
		</input>
		<output>
            <integer name="schedule_size" />
          <for count="schedule_size" from="0" name="i">
                <integer array_index="i" name="id" />
                <integer array_index="i" name="create_date" />
                <integer array_index="i" name="change_date" />
                <integer name="items_size" />
              <for count="items_size" from="0" name="j">
                <integer array_index="i,j" name="scheme_id" />
                <integer array_index="i,j" name="start_date" />
                <integer array_index="i,j" name="end_date" />
                <double array_index="i,j" name="coeff" />
              </for>
          </for>
		</output>
	</function>


*/
