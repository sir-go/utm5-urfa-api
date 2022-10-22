package fn

// rpcf_get_traffic_collectors_list
func x5064(c conn, _ Dict) Dict {
	c.Call(0x5064)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5064" name="rpcf_get_traffic_collectors_list">
    <input />
    <output>
            <integer name="size" />
            <for count="size" from="0" name="i">
              <integer array_index="i" name="id" />
              <string array_index="i" name="name" />
              <string array_index="i" name="comments" />
              <integer array_index="i" name="state" />
            </for>
    </output>
  </function>


*/
