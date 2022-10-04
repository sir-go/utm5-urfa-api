package fn

// rpcf_get_hotspot_networks
func x10201(c conn, p Dict) Dict {
	c.Call(0x10201)
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"networks": getMapOf(c, func() Dict {
			return Dict{
				"ip":   c.GetI(),
				"mask": c.GetI(),
			}
		}),
	}
}

/* xml:
<function id="0x10201" name="rpcf_get_hotspot_networks">
        <input>
            <integer name="service_id" />
        </input>
        <output>
            <integer name="cnt" />
            <for count="cnt" from="0" name="i">
                <integer array_index="i" name="ip" />
                <integer array_index="i" name="mask" />
            </for>
        </output>
    </function>


*/
