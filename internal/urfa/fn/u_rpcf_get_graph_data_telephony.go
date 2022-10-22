package fn

// rpcf_get_graph_data_telephony
func x5093(c conn, p Dict) Dict {
	c.Call(0x5093)
	putI(c, p, "uid")
	putI(c, p, "aid")
	putI(c, p, "t_start")
	putI(c, p, "t_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5093" name="rpcf_get_graph_data_telephony">
    <input>
      <integer name="uid" />
      <integer name="aid" />
      <integer name="t_start" />
      <integer name="t_end" />
    </input>
    <output>
      <integer name="gdata_size" />
      <for count="gdata_size" from="0" name="i">
        <integer array_index="i" name="date" />
        <integer name="size" />
        <set dst="size_array" dst_index="i" src="size" />
        <for count="size" from="0" name="j">
          <string array_index="i,j" name="name" />
          <double array_index="i,j" name="bytes" />
        </for>
      </for>
    </output>
  </function>


*/
