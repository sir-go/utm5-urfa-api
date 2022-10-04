package fn

// rpcf_dealer_get_tclasses
func x7000002d(c conn, _ Dict) Dict {
	c.Call(0x7000002d)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000002d" name="rpcf_dealer_get_tclasses">
      <input>
      </input>
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
