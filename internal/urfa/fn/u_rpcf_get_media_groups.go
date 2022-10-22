package fn

// rpcf_get_media_groups
func x450e(c conn, _ Dict) Dict {
	c.Call(0x450e)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x450e" name="rpcf_get_media_groups">
      <output>
        <integer name="size" />
        <for count="size" from="0" name="i">
          <long array_index="i" name="id" />
          <string array_index="i" name="name" />
          <integer array_index="i" name="type" />
        </for>
      </output>
    </function>


*/
