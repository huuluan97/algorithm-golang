package concurrency

import (
	"crypto/md5"
	"errors"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	c := make(chan result)
	errc := make(chan error, 1)

	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			wg.Add(1)

			go func() { // HL
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}: // HL
				case <-done: // HL
				}
				wg.Done()
			}()

			select {
			case <-done: // HL
				return errors.New("walk canceled")
			default:
				return nil
			}

		})
		go func() { // HL
			wg.Wait()
			close(c) // HL
		}()
		errc <- err // HL
	}()

	return c, errc
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)

	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	return m, nil
}

//func onHandleParallelismConcurrency() {
//	// Calculate the MD5 sum of all files under the specified directory,
//	// then print the results sorted by path name.
//	m, err := MD5All(os.Args[1])
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	var paths []string
//	for path := range m {
//		paths = append(paths, path)
//	}
//	sort.Strings(paths)
//	for _, path := range paths {
//		fmt.Printf("%x  %s\n", m[path], path)
//	}
//}