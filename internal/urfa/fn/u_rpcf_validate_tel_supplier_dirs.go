package fn

// rpcf_validate_tel_supplier_dirs
func x10321(c conn, p Dict) Dict {
	c.Call(0x10321)
	putI(c, p, "supplier_id")
	forEach(c, p, "dirs", func(item interface{}) {
		c.PutInt(item.(int))
	})
	c.Send()

	return Dict{
		"result":              c.GetI(),
		"founded_dir_id":      c.GetI(),
		"founded_zone_id":     c.GetI(),
		"founded_supplier_id": c.GetI(),
	}
}

/* xml:
<function id="0x10321" name="rpcf_validate_tel_supplier_dirs">
    <input>
      <integer name="supplier_id" />
      <integer name="dirs_cnt" />
      <for count="dirs_cnt" from="0" name="i">
          <integer array_index="i" name="dir" />
      </for>
    </input>
    <output>
      <integer name="result" />
      <integer name="founded_dir_id" />
      <integer name="founded_zone_id" />
      <integer name="founded_supplier_id" />
    </output>
  </function>


*/
