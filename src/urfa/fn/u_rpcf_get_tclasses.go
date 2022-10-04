package fn

// rpcf_get_tclasses
func x2300(c conn, _ Dict) Dict {
	c.Call(0x2300)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2300" name="rpcf_get_tclasses">
    <input />
    <output>
      <integer name="tclass_list_size" />
      <for count="tclass_list_size" from="0" name="i">
        <integer array_index="i" name="tclass_id" />
        <string array_index="i" name="tclass_name" />
        <integer array_index="i" name="graph_color" />
        <integer array_index="i" name="is_display" />
        <integer array_index="i" name="is_fill" />
      </for>
    </output>
  </function>


*/
