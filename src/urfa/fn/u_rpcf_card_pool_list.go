package fn

// rpcf_card_pool_list
func x4201(c conn, _ Dict) Dict {
	c.Call(0x4201)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4201" name="rpcf_card_pool_list">
    <input />
    <output>
      <integer name="info_size" />
      <for count="info_size" from="0" name="i">
        <integer array_index="i" name="pool_id" />
        <integer array_index="i" name="cards" />
        <integer array_index="i" name="cards_used" />
        <integer array_index="i" name="first_update" />
        <integer array_index="i" name="last_update" />
      </for>
    </output>
  </function>


*/
