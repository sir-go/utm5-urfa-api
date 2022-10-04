package fn

// rpcf_get_fwrules_list_new
func x5020(c conn, _ Dict) Dict {
	c.Call(0x5020)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5020" name="rpcf_get_fwrules_list_new">
    <input />
    <output>
      <integer name="rules_count" />
      <for count="rules_count" from="0" name="i">
        <integer array_index="i" name="rule_id" />
	<integer array_index="i" name="flags" />
	<long array_index="i" name="events" />
	<integer array_index="i" name="router_id" />
	<integer array_index="i" name="tariff_id" />
	<integer array_index="i" name="group_id" />
	<integer array_index="i" name="user_id" />
	<string array_index="i" name="rule" />
	<string array_index="i" name="comment" />
      </for>
   </output>
  </function>


*/
