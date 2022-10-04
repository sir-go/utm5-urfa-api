package fn

// rpcf_get_suppliers_list
func x8014(c conn, _ Dict) Dict {
	c.Call(0x8014)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x8014" name="rpcf_get_suppliers_list">
    <output>
      <integer name="size" />
      <for count="size" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="name" />
        <string array_index="i" name="short_name" />
        <string array_index="i" name="jur_adress" />
        <string array_index="i" name="act_adress" />
        <string array_index="i" name="inn" />
        <string array_index="i" name="kpp" />
        <integer array_index="i" name="bank_id" />
        <string array_index="i" name="account" />
        <string array_index="i" name="headman" />
        <string array_index="i" name="bookeeper" />
        <string array_index="i" name="short_headman" />
        <string array_index="i" name="short_bookeeper" />
        <double array_index="i" name="supp_balance" />
        <double array_index="i" name="tax_rate" />
      </for>
    </output>
  </function>


*/
