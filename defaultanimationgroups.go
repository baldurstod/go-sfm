package sfm

import (
	"fmt"
	"io/fs"

	"github.com/baldurstod/vdf"
)

type group struct {
	Color
	root []string
}

var groups map[string]group

var unknownGroup = group{root: []string{"Unknown"}, Color: [4]uint8{0, 128, 255, 255}}

func GetAnimationGroup(name string) group {
	if groups == nil {
		parseAnimationGroups()
	}

	group, ok := groups[name]
	if !ok {
		return unknownGroup
	}

	return group
}

func parseAnimationGroups() error {
	if groups == nil {
		groups = make(map[string]group)
	}

	var resourcesFs = &Resources

	buf, err := fs.ReadFile(fs.FS(resourcesFs), "resources/sfm_default_animation_groups.vcfg")
	if err != nil {
		return fmt.Errorf("failed to open resources/: <%w>", err)
	}

	vdf := vdf.VDF{}
	root := vdf.Parse(buf)

	err = initAnimationGroups(root)
	if err != nil {
		return fmt.Errorf("failed to init groups: <%w>", err)
	}

	return nil
}

func initAnimationGroups(root vdf.KeyValue) error {
	groupFile, err := root.Get("groupFile")
	if err != nil {
		return err
	}

	for _, val := range groupFile.GetChilds() {
		//log.Println(val)
		initAnimationGroup(val, make([]string, 0, 5))
	}

	return nil
}

func initAnimationGroup(group *vdf.KeyValue, st []string) error {
	dst := make([]string, len(st), cap(st))
	copy(dst, st)
	dst = append(dst, group.Key)

	controls, _ := group.GetAll("control")
	if len(controls) != 0 {
		return addAnimationGroup(group, dst)
	}

	for _, val := range group.GetChilds() {
		err := initAnimationGroup(val, dst)
		if err != nil {
			return err
		}
	}

	return nil
}

func addAnimationGroup(gr *vdf.KeyValue, st []string) error {
	var r, g, b, a uint8
	if color, err := gr.GetString("groupColor"); err != nil {
		_, err := fmt.Sscanf(color, "%d %d %d %d", &r, &g, &b, &a)
		if err != nil {
			r = 255
			g = 255
			b = 255
			a = 255
		}
	} else {
		r = 0
		g = 128
		b = 255
		a = 255
	}

	controls, _ := gr.GetAll("control")

	for _, val := range controls {
		s, err := val.ToString()
		if err != nil {
			return err
		}
		groups[s] = group{root: st, Color: [4]uint8{r, g, b, a}}
	}
	return nil
}
