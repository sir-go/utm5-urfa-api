package fn

// rpcf_add_coeff_scheme
func x4702(c conn, p Dict) Dict {
	c.Call(0x4702)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"id": c.GetI(),
	}
}

/* xml:
<function id="0x4702" name="rpcf_add_coeff_scheme">
		<input>
            <integer name="id" />
            <string name="name" />
            <string name="comment" />
            <integer name="create_date" />
            <integer name="change_date" />
            <integer name="items_size" />
          <for count="items_size" from="0" name="i">
            <integer array_index="i" name="scheme_id" />
            <integer array_index="i" name="delay" />
            <integer array_index="i" name="duration" />
            <double array_index="i" name="coeff" />
          </for>
		</input>
		<output>
            <integer name="id" />
		</output>
	</function>


*/
