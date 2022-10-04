package fn

// rpcf_user5_get_unused_prepaid
func xU5502(c conn, _ Dict) Dict {
	c.Call(-0x5502)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x5502" name="rpcf_user5_get_unused_prepaid">
    <input />
    <output>
      <integer name="bytes_in_mbyte" />
      <integer name="links_size" />
      <for count="links_size" from="0" name="i">
        <integer name="pinfo_size" />
        <set dst="pinfo_size_array" dst_index="i" src="pinfo_size" />
        <for count="pinfo_size" from="0" name="j">
          <integer array_index="i,j" name="id" />
          <long array_index="i,j" name="value" />
        </for>
      </for>
    </output>
  </function>


*/
