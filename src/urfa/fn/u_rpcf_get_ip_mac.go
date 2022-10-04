package fn

// rpcf_get_ip_mac
func x2814(c conn, _ Dict) Dict {
	c.Call(0x2814)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2814" name="rpcf_get_ip_mac">
    <input />
    <output>
      <integer name="ip_mac_size" />
      <for count="ip_mac_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="name" />
      </for>
    </output>
  </function>


*/
