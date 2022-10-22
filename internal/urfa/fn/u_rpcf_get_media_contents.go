package fn

// rpcf_get_media_contents
func x450d(c conn, _ Dict) Dict {
	c.Call(0x450d)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x450d" name="rpcf_get_media_contents">
      <output>
        <integer name="size" />
        <for count="size" from="0" name="i">
          <integer array_index="i" name="id" />
          <string array_index="i" name="name" />
          <integer array_index="i" name="type" />
        </for>
      </output>
    </function>


*/
