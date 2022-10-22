package fn

// rpcf_set_hotspot_networks
func x10202(c conn, p Dict) Dict {
	c.Call(0x10202)
	putI(c, p, "service_id")
	forEachDict(c, p, "networks", func(ai Dict) {
		putI(c, ai, "ip")
		putI(c, ai, "mask")
	})
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x10202" name="rpcf_set_hotspot_networks">
        <input>
            <integer name="service_id" />
            <integer name="cnt" />
            <for count="cnt" from="0" name="i">
                <integer array_index="i" name="ip" />
                <integer array_index="i" name="mask" />
            </for>
        </input>
        <output>
            <integer name="result" />
        </output>
    </function>


*/
