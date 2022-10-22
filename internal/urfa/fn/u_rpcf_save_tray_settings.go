package fn

// rpcf_save_tray_settings
func x212b(c conn, p Dict) Dict {
	c.Call(0x212b)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x212b" name="rpcf_save_tray_settings">
    <input>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <string array_index="i" name="attribute" />
	<string array_index="i" name="value" />
      </for>
    </input>
    <output>
      <integer name="result" />
    </output>
  </function>


*/
