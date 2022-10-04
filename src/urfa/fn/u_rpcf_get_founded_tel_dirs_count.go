package fn

// rpcf_get_founded_tel_dirs_count
func x1031f(c conn, p Dict) Dict {
	c.Call(0x1031f)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"count": c.GetI(),
	}
}

/* xml:
<function id="0x1031f" name="rpcf_get_founded_tel_dirs_count">
      <input>
          <integer name="status" />
          <integer name="params_cnt" />
          <for count="params_cnt" from="0" name="i">
              <integer array_index="i" name="field" />
              <integer array_index="i" name="criteria" />
              <string array_index="i" name="value" />
          </for>
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
