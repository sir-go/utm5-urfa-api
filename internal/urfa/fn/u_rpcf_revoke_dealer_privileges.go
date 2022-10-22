package fn

// rpcf_revoke_dealer_privileges
func x1301c(c conn, p Dict) Dict {
	c.Call(0x1301c)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function comment="Revoke privileges to system entities" id="0x1301c" name="rpcf_revoke_dealer_privileges">
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
