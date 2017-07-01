// go doc filepath | awk -F'[ \(]' 'BEGIN{ print "package main; import \"fmt\"; import \"path/filepath\"; func main(){"} /func.*\).*string/{ printf "fmt.Printf(\"called %%q: %%q\\n\",\"%s\",fmt.Sprint(filepath.%s(`/windows/system32/drivers/etc/hosts`)))\n", $2, $2} END{ print "};" }' | gofmt
package main

import "fmt"
import "path/filepath"
import "time"

func main() {
	fmt.Println(time.Now())
	fmt.Printf("called %q: %q\n", "Abs", fmt.Sprint(filepath.Abs(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "Base", fmt.Sprint(filepath.Base(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "Base", fmt.Sprint(filepath.Base(`/var/lib/docker/containers/*/*-json.log`)))
	fmt.Printf("called %q: %q\n", "Dir", fmt.Sprint(filepath.Dir(`/var/lib/docker/containers/*/*-json.log`)))
	fmt.Printf("called %q: %q\n", "Clean", fmt.Sprint(filepath.Clean(`/var/lib/docker/containers/*/*-json.log`)))
	fmt.Printf("called %q: %q\n", "Base2“”‘【filefiLastLogosfu", fmt.Sprint(filepath.Base(`/windows/system32/drivers/etc/hosts/`)))
	fmt.Printf("called %q: %q\n", "Dir", fmt.Sprint(filepath.Dir(`/windows/system32/drivers/etc/hosts/`)))
	fmt.Printf("called %q: %q\n", "Dir", fmt.Sprint(filepath.Dir(`/windows/system32/drivers/etc/hosts/afdafd`)))
	fmt.Printf("called %q: %q\n", "EvalSymlinks", fmt.Sprint(filepath.EvalSymlinks(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "Ext", fmt.Sprint(filepath.Ext(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "FromSlash", fmt.Sprint(filepath.FromSlash(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "Glob", fmt.Sprint(filepath.Glob(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "Join", fmt.Sprint(filepath.Join(`/windows/system32/drivers/etc/hosts`)))
	// Rel needs second arg
	fmt.Printf("called %q: %q\n", "Rel", fmt.Sprint(filepath.Rel(`/windows/system32/drivers/etc/hosts`, `/windows/`)))
	fmt.Printf("called %q: %q\n", "Split", fmt.Sprint(filepath.Split(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "SplitList", fmt.Sprint(filepath.SplitList(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "ToSlash", fmt.Sprint(filepath.ToSlash(`/windows/system32/drivers/etc/hosts`)))
	fmt.Printf("called %q: %q\n", "VolumeName", fmt.Sprint(filepath.VolumeName(`/windows/system32/drivers/etc/hosts`)))
	fmt.Println("abcdfe"[2:5])
	arr2d := make([][]string, 4)
	fmt.Println(len(arr2d), cap(arr2d))
	arr2d[0] = make([]string, 3)
	fmt.Println(len(arr2d[0]), cap(arr2d[0]))
	arr2d[1] = make([]string, 10)
	fmt.Println(len(arr2d[1]), cap(arr2d[1]))
	var aa []int
	bb := []int{3, 4}
	fmt.Println(aa)
	aa = append(aa, bb...)
	fmt.Println(aa)
}
