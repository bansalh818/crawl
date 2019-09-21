package pool

type Job struct {  
    id       int
    Page int
}
type Result struct {  
    job         Job
    sumofdigits int
}




// //
// type Work struct {
// 	ID  int
// 	Job string
// }

// type Worker struct {
// 	ID            int
// 	WorkerChannel chan chan Work
// 	Channel       chan Work
// 	End           chan bool
// }

// // Start worker
// func (w *Worker) Start() {
// 	go func() {
// 		for {
// 			w.WorkerChannel <- w.Channel
// 			select {
// 			case job := <-w.Channel:
// 				// do work
				
// 			case <-w.End:
// 				return
// 			}
// 		}
// 	}()
// }

// // end worker
// func (w *Worker) Stop() {
// 	w.End <- true
// }
