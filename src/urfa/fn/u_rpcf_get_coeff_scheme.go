package fn

// rpcf_get_coeff_scheme
func x4700(c conn, p Dict) Dict {
	c.Call(0x4700)
	putI(c, p, "scheme_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4700" name="rpcf_get_coeff_scheme">
		<input>
            <integer name="scheme_id" />
		</input>
		<output>
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
		</output>
	</function>


*/
