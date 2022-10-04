package fn

// rpcf_get_switch_types_list
func x508(c conn, _ Dict) Dict {
	c.Call(0x508)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x508" name="rpcf_get_switch_types_list">
        <output>
            <integer name="size" />
            <for count="size" from="0" name="i">
                <integer array_index="i" name="id" />
                <string array_index="i" name="name" />
                <string array_index="i" name="supp_volumes" />
                <integer array_index="i" name="port_start_offset" />
                <integer array_index="i" name="tec_id_type" />
                <integer array_index="i" name="tec_id_len" />
                <integer array_index="i" name="tec_id_disp" />
                <integer array_index="i" name="tec_id_offset" />
                <integer array_index="i" name="port_id_type" />
                <integer array_index="i" name="port_id_len" />
                <integer array_index="i" name="port_id_disp" />
                <integer array_index="i" name="port_id_offset" />
                <integer array_index="i" name="vlan_id_type" />
                <integer array_index="i" name="vlan_id_len" />
                <integer array_index="i" name="vlan_id_disp" />
                <integer array_index="i" name="vlan_id_offset" />
            </for>
        </output>
    </function>


*/
