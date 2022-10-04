package fn

// rpcf_grant_privileges_to_dealer
func x1301a(c conn, p Dict) Dict {
	c.Call(0x1301a)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function comment="Grant privileges to dealer on system entities" id="0x1301a" name="rpcf_grant_privileges_to_dealer">
      <input>
        <integer name="dealer_id" />
	<integer name="entity_type" />
	<integer default="size(entity_id_array)" name="entity_id_count" />
	<for count="size(entity_id_array)" from="0" name="i">
	    <integer array_index="i" name="entity_id_array" />
	</for>
      </input>
      <output>
        <integer name="result" />
      </output>
    </function>


*/
