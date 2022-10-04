package fn

// rpcf_get_tel_directions_count
func x1031e(c conn, p Dict) Dict {
	c.Call(0x1031e)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"count": c.GetI(),
	}
}

/* xml:
<function id="0x1031e" name="rpcf_get_tel_directions_count">
      <input>
        <integer name="skip_ids_cnt" />
        <for count="skip_ids_cnt" from="0" name="i">
            <integer array_index="i" name="skip_id" />
        </for>
      </input>
      <output>
          <integer name="count" />
      </output>
  </function>


*/
