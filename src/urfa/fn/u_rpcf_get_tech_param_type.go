package fn

// rpcf_get_tech_param_type
func x9002(c conn, _ Dict) Dict {
	c.Call(0x9002)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x9002" name="rpcf_get_tech_param_type">
    <input />
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="name" />
      </for>
    </output>
  </function>


*/
