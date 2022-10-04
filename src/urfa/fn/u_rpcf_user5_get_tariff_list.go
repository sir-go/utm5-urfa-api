package fn

// rpcf_user5_get_tariff_list
func xU401d(c conn, _ Dict) Dict {
	c.Call(-0x401d)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x401d" name="rpcf_user5_get_tariff_list">
    <output>
      <integer name="size" />
      <for count="size" from="0" name="i">
        <integer array_index="i" name="tp_id" />
      </for>
    </output>
  </function>


*/
