package fn

// rpcf_get_suppliers_zones
func x8019(c conn, _ Dict) Dict {
	c.Call(0x8019)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x8019" name="rpcf_get_suppliers_zones">
    <output>
      <integer name="ret_code" />
      <integer name="size" />
      <for count="size" from="0" name="i">
        <integer array_index="i" name="zone_id" />
        <integer array_index="i" name="supplier_id" />
      </for>
    </output>
  </function>


*/
