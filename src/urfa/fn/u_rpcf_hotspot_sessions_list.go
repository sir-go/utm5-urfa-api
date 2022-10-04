package fn

// rpcf_hotspot_sessions_list
func x4308(c conn, _ Dict) Dict {
	c.Call(0x4308)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4308" name="rpcf_hotspot_sessions_list">
        <input />
        <output>
            <integer name="sess_count" />
            <for count="sess_count" from="0" name="i">
                <integer array_index="i" name="dhs_sess_id" />
                <string array_index="i" name="nas_sess_id" />
                <string array_index="i" name="web_sess_id" />
                <integer array_index="i" name="start_date" />
                <integer array_index="i" name="end_date" />
                <integer array_index="i" name="slink_id" />
                <integer array_index="i" name="ip" />
                <string array_index="i" name="login" />
                <double array_index="i" name="banalce" />
            </for>
        </output>
    </function>


*/
