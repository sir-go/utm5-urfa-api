package fn

// rpcf_get_tel_suppliers_directions
func x8023(c conn, _ Dict) Dict {
	c.Call(0x8023)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x8023" name="rpcf_get_tel_suppliers_directions">
    <input />
    <output>
      <integer name="dirs_cnt" />
      <for count="dirs_cnt" from="0" name="i">
        <integer array_index="i" name="dir_id" />
        <integer array_index="i" name="supplier_id" />
      </for>
    </output>
  </function>


*/
