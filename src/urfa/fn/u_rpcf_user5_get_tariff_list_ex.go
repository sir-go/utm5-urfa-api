package fn

// rpcf_user5_get_tariff_list_ex
func xU401f(c conn, _ Dict) Dict {
	c.Call(-0x401f)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x401f" name="rpcf_user5_get_tariff_list_ex">
    <output>
      <integer name="size" />
      <for count="size" from="0" name="i">
        <integer array_index="i" name="tp_link_id" />
        <integer array_index="i" name="tp_id" />
      </for>
    </output>
  </function>


*/
