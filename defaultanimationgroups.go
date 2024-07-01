package sfm

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/baldurstod/go-vector"
	"github.com/baldurstod/vdf"
)

type group struct {
	Name  string
	Color vector.Vector4[uint8]
	root  []string
}

var groups map[string]*group

func GetAnimationGroup(name string) *group {
	if groups == nil {
		parseAnimationGroups()
	}

	return groups[name]
}

func parseAnimationGroups() error {
	if groups == nil {
		groups = make(map[string]*group)
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
	groupFile, ok := root.Get("groupFile")
	if !ok {
		return errors.New("unable to find groupFile")
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
	if color, ok := gr.GetString("groupColor"); ok {
		_, err := fmt.Sscanf(color, "%d %d %d %d", &r, &g, &b, &a)
		if err != nil {
			r = 0
			g = 0
			b = 0
			a = 255
		}
	}

	controls, _ := gr.GetAll("control")

	for _, val := range controls {
		s, ok := val.ToString()
		if !ok {
			return errors.New("value is not of type string")
		}
		groups[s] = &group{root: st}
		groups[s].Color.Set(r, g, b, a)
	}
	return nil
}
