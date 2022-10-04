package fn

// rpcf_get_session_list
func x2123(c conn, _ Dict) Dict {
	c.Call(0x2123)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2123" name="rpcf_get_session_list">
      <input />
      <output>
        <integer name="count" />
	  <for count="count" from="0" name="i">
              <integer array_index="i" name="sid" />
              <integer array_index="i" name="type" />
              <integer array_index="i" name="staff_id" />
              <string array_index="i" name="host" />
              <integer array_index="i" name="port" />

	  </for>
      </output>
  </function>


*/
