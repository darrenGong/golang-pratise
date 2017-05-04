package main

import "fmt"

func main() {
	fmt.Println(fmt.Sprintf("hello %s", "world", "xiaotao"))
	sql := fmt.Sprintf("SELECT private.id, private.ip, private.mac, private.subnetwork_id, private.nat_no from t_ip_object object, t_ip_private private WHERE object.subnetwork_id is not null AND private.ip = object.ip AND object.object_id in(SELECT gw.object_id FROM t_cnat_gw gw WHERE gw.gw_id = '%s' AND gw.account_id = %d)",
					   "Xiaotao", 100000001)
	fmt.Println(sql)
}
