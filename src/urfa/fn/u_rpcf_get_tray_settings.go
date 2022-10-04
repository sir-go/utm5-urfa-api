package fn

// rpcf_get_tray_settings
func x212a(c conn, _ Dict) Dict {
	c.Call(0x212a)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x212a" name="rpcf_get_tray_settings">
    <input />
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <string array_index="i" name="attribute" />
	<string array_index="i" name="value" />
      </for>
    </output>
  </function>


*/
