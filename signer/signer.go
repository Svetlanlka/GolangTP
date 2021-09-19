package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func ExecutePipeline(workers ...job) {
	var in chan interface{}
	wg := new(sync.WaitGroup)

	for _, work := range workers {
		out := make(chan interface{}, 1)

		wg.Add(1)
		go func(wg *sync.WaitGroup, work job, in chan interface{}, out chan interface{}) {
			defer wg.Done()
			work(in, out)
			close(out)
		}(wg, work, in, out)

		in = out
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	wgComplex := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for dataRaw := range in {
		wgComplex.Add(1)
		go func(dataRaw interface{}, wgComplex *sync.WaitGroup) {
			defer wgComplex.Done()
			data := fmt.Sprint(dataRaw)
			data1 := make(chan string, 1)
			data2 := make(chan string, 1)

			mu.Lock()
			dataMd5 := DataSignerMd5(data)
			mu.Unlock()

			go func(data string) {
				data1 <- DataSignerCrc32(data)
			}(data)
			go func(dataMd5 string) {
				data2 <- DataSignerCrc32(dataMd5)
			}(dataMd5)

			hash := <-data1 + "~" + <-data2
			out <- hash
		}(dataRaw, wgComplex)
	}
	wgComplex.Wait()
}

func MultiHash(in, out chan interface{}) {
	wgComplex := new(sync.WaitGroup)
	for dataRaw := range in {
		wgComplex.Add(1)
		go func(dataRaw interface{}, wgComplex *sync.WaitGroup) {
			defer wgComplex.Done()
			var hash sync.Map
			var result string = ""
			data, ok := dataRaw.(string)
			if !ok {
				fmt.Println("cant convert input data to string")
				return
			}
			wg := new(sync.WaitGroup)

			for th := 0; th < 6; th++ {
				wg.Add(1)
				go func(th int, data string, hash *sync.Map, wg *sync.WaitGroup) {
					defer wg.Done()
					hash.Store(th, DataSignerCrc32(strconv.Itoa(th)+data))
				}(th, data, &hash, wg)
			}
			wg.Wait()

			for th := 0; th < 6; th++ {
				h, ok := hash.Load(th)
				if !ok {
					fmt.Println("cant convert input data to string")
					return
				}
				strHash := fmt.Sprint(h)
				result += strHash
			}

			out <- result
		}(dataRaw, wgComplex)
	}
	wgComplex.Wait()
}

func CombineResults(in, out chan interface{}) {
	var result string = ""
	var hash []string

	for dataRaw := range in {
		data, ok := dataRaw.(string)
		if !ok {
			fmt.Println("cant convert input data to string")
			return
		}

		hash = append(hash, data)
	}

	sort.Strings(hash)

	for i := range hash {
		result += hash[i]
		if i < len(hash)-1 {
			result += "_"
		}
	}
	out <- result
}
