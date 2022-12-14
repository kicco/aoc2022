package main

type Fs struct {
	Root *Node
	Path []*Node
}

func (fs *Fs) Size() int {
	return fs.Path[len(fs.Path)-1].Size()
}

func (fs *Fs) Cd(dir string) {

	if dir == "/" {
		fs.Path = nil
		fs.Path = append(fs.Path, fs.Root)
		return
	}

	if dir == ".." {
		if len(fs.Path) > 1 {
			fs.Path = fs.Path[:len(fs.Path)-1]
		}
		return
	}

	for _, child := range fs.Path[len(fs.Path)-1].Sub {
		if child.Name == dir {
			fs.Path = append(fs.Path, child)
			return
		}
	}

	panic("Path not found")
}

func (fs *Fs) Mkdir(name string) {
	fs.Path[len(fs.Path)-1].Mkdir(name)
}

func (fs *Fs) Touch(name string, size int) {
	fs.Path[len(fs.Path)-1].Touch(name, size)
}
