package fn

// rpcf_get_dealer_privileges
func x1301b(c conn, p Dict) Dict {
	c.Call(0x1301b)
	putI(c, p, "dealer_id")
	putI(c, p, "entity_type")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Returns entity ID list granted for dealer" id="0x1301b" name="rpcf_get_dealer_privileges">
      <input>
        <integer name="dealer_id" />
	<integer name="entity_type" />
      </input>
      <output>
	<integer name="entity_id_count" />
	<for count="entity_id_count" from="0" name="i">
	    <integer array_index="i" name="entity_id_array" />
	</for>
      </output>
    </function>


*/
