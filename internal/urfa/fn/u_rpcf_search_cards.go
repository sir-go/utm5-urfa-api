package fn

// rpcf_search_cards
func x1201(c conn, p Dict) Dict {
	c.Call(0x1201)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1201" name="rpcf_search_cards">
    <input>
      <integer name="select_type" />
      <integer default="size(what_id)" name="patterns_count" />
      <for count="size(what_id)" from="0" name="i">
        <integer array_index="i" name="what_id" />
        <integer array_index="i" name="criteria_id" />
        <string array_index="i" name="pattern" />
      </for>
    </input>
    <output>
      <integer name="cards_size" />
      <for count="cards_size" from="0" name="i">
        <integer array_index="i" name="card_id" />
        <integer array_index="i" name="pool_id" />
        <string array_index="i" name="secret" />
        <double array_index="i" name="balance" />
        <integer array_index="i" name="currency" />
        <integer array_index="i" name="expire" />
        <integer array_index="i" name="days" />
        <integer array_index="i" name="is_used" />
        <integer array_index="i" name="tp_id" />
      </for>
    </output>
  </function>


*/
