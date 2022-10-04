package fn

// rpcf_dealer_get_charge_policy_list
func x70000066(c conn, _ Dict) Dict {
	c.Call(0x70000066)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000066" name="rpcf_dealer_get_charge_policy_list">
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
