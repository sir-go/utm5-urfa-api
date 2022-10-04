package fn

// rpcf_get_charge_policy_list
func x15102(c conn, _ Dict) Dict {
	c.Call(0x15102)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x15102" name="rpcf_get_charge_policy_list">
    <input />
    <output>
      <integer name="policy_count" />
      <for count="policy_count" from="0" name="i">
         <integer array_index="i" name="policy_id_array" />
         <integer array_index="i" name="flags_array" />
         <string array_index="i" name="name_array" />
         <integer array_index="i" name="tm_count" />
         <for count="tm_count" from="0" name="e">
            <integer array_index="i,e" name="timemark" />
         </for>
      </for>
    </output>
  </function>


*/
