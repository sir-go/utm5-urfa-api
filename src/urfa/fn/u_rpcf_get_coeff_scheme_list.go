package fn

// rpcf_get_coeff_scheme_list
func x4701(c conn, _ Dict) Dict {
	c.Call(0x4701)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4701" name="rpcf_get_coeff_scheme_list">
		<input>
		</input>
		<output>
            <integer name="scheme_size" />
          <for count="scheme_size" from="0" name="i">
                <integer array_index="i" name="id" />
                <string array_index="i" name="name" />
                <string array_index="i" name="comment" />
                <integer array_index="i" name="create_date" />
                <integer array_index="i" name="change_date" />
                <integer name="items_size" />
              <for count="items_size" from="0" name="j">
                <integer array_index="i,j" name="scheme_id" />
                <integer array_index="i,j" name="delay" />
                <integer array_index="i,j" name="duration" />
                <double array_index="i,j" name="coeff" />
              </for>
          </for>
		</output>
	</function>


*/
