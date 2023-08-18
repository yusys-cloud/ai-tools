// Author: yangzq80@gmail.com
// Date: 2023/6/15
package flow

import "flag"

func main() {
	path := flag.String("path", "conf.json", "--path=conf.json")
	flag.Parse()
	New(*path).Run()
}
